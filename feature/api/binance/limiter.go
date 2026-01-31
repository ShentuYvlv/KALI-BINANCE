package binance

import (
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/adshao/go-binance/v2/common"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
)

type rateLimiter struct {
	mu         sync.Mutex
	capacity   float64
	tokens     float64
	refillRate float64
	lastRefill time.Time
	banUntil   time.Time
}

func newRateLimiter(perMin int) *rateLimiter {
	if perMin <= 0 {
		perMin = 1800
	}
	capacity := float64(perMin)
	refillRate := capacity / 60.0
	return &rateLimiter{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

func (l *rateLimiter) setRate(perMin int) {
	if perMin <= 0 {
		perMin = 1800
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	capacity := float64(perMin)
	l.capacity = capacity
	l.refillRate = capacity / 60.0
	if l.tokens > l.capacity {
		l.tokens = l.capacity
	}
}

func (l *rateLimiter) setBanUntil(t time.Time) {
	if t.IsZero() {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if t.After(l.banUntil) {
		l.banUntil = t
	}
}

func (l *rateLimiter) acquire(weight int) {
	if weight <= 0 {
		weight = 1
	}
	for {
		var wait time.Duration
		now := time.Now()
		l.mu.Lock()
		if now.Before(l.banUntil) {
			wait = l.banUntil.Sub(now)
			l.mu.Unlock()
			time.Sleep(wait)
			continue
		}
		elapsed := now.Sub(l.lastRefill).Seconds()
		if elapsed > 0 {
			l.tokens = min(l.capacity, l.tokens+elapsed*l.refillRate)
			l.lastRefill = now
		}
		if l.tokens >= float64(weight) {
			l.tokens -= float64(weight)
			l.mu.Unlock()
			return
		}
		deficit := float64(weight) - l.tokens
		wait = time.Duration(deficit/l.refillRate*float64(time.Second)) + 50*time.Millisecond
		l.mu.Unlock()
		time.Sleep(wait)
	}
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

var limiter = newRateLimiter(loadRateLimit())

func loadRateLimit() int {
	if v, err := config.Int("binance::rate_limit_per_min"); err == nil && v > 0 {
		return v
	}
	return 1800
}

func refreshRateLimit() {
	limiter.setRate(loadRateLimit())
}

func withLimiter(weight int, fn func() error) error {
	refreshRateLimit()
	limiter.acquire(weight)
	if fn == nil {
		return nil
	}
	err := fn()
	if err != nil {
		handleRateLimitError(err)
	}
	return err
}

func handleRateLimitError(err error) {
	if err == nil {
		return
	}
	var apiErr *common.APIError
	if strings.Contains(err.Error(), "banned until") || strings.Contains(err.Error(), "Way too many requests") {
		if t := parseBanUntil(err.Error()); !t.IsZero() {
			logs.Error("binance rate limit ban until:", t)
			limiter.setBanUntil(t)
		}
		return
	}
	if strings.Contains(err.Error(), "code=-1003") {
		if t := parseBanUntil(err.Error()); !t.IsZero() {
			logs.Error("binance rate limit ban until:", t)
			limiter.setBanUntil(t)
		}
		return
	}
	if ok := strings.HasPrefix(err.Error(), "<APIError>"); ok {
		if strings.Contains(err.Error(), "banned until") {
			if t := parseBanUntil(err.Error()); !t.IsZero() {
				logs.Error("binance rate limit ban until:", t)
				limiter.setBanUntil(t)
			}
		}
	}
	if errorsAs(err, &apiErr) {
		if apiErr.Code == -1003 {
			if t := parseBanUntil(err.Error()); !t.IsZero() {
				logs.Error("binance rate limit ban until:", t)
				limiter.setBanUntil(t)
			}
		}
	}
}

func errorsAs(err error, target **common.APIError) bool {
	if err == nil {
		return false
	}
	if apiErr, ok := err.(*common.APIError); ok {
		*target = apiErr
		return true
	}
	return false
}

func parseBanUntil(msg string) time.Time {
	re := regexp.MustCompile(`banned until\s+(\d+)`)
	matches := re.FindStringSubmatch(msg)
	if len(matches) < 2 {
		return time.Time{}
	}
	ms, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return time.Time{}
	}
	return time.Unix(0, ms*int64(time.Millisecond))
}
