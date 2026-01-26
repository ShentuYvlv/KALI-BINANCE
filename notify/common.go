package notify

import (
	"go_binance_futures/models"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
)

func nowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func getStatusColor(status string) string {
	var color string
	switch status {
		case "success":
			color = "#008000"
		case "fail":
			color = "#FF0000"
		default:
			color = "#008000"
	}
	return color
}

func GetNotifyChannel() (pusher Pusher) {
	var notification_channel, _ = config.String("notification::channel") // 通知方式
	switch (notification_channel) {
		case "dingding":
			pusher = DingDing{}
		case "slack":
			pusher = Slack{}
		default:
			pusher = DingDing{}
	}
	return pusher
}

func GetNotifyConfig(pusher Pusher) (notifyConfig models.NotifyConfig) {
	var notification_channel, _ = config.String("notification::channel") // 通知方式
	moduleName := pusher.GetModuleName()
	if moduleName == "" {
		return notifyConfig
	}
	o := orm.NewOrm()
	o.QueryTable("notify_config").
		Filter("module", moduleName).
		Filter("channel", notification_channel).
		Filter("enable", 1).
		OrderBy("-id").
		One(&notifyConfig)

	return notifyConfig
}

func SaveNotification(channel string, pusher Pusher, content string, level ...string) {
	moduleName := pusher.GetModuleName()
	if moduleName == "" {
		moduleName = "system"
	}
	levelValue := ""
	if len(level) > 0 {
		levelValue = level[0]
	}
	if levelValue == "" {
		levelValue = detectLevel(moduleName, content)
	}
	title := extractTitle(content)
	message := models.NotifyMessage{
		Module: moduleName,
		Channel: channel,
		Level: levelValue,
		Title: title,
		Content: content,
		CreateTime: time.Now().Unix(),
	}
	o := orm.NewOrm()
	_, _ = o.Insert(&message)
	cutoff := time.Now().Add(-30 * 24 * time.Hour).Unix()
	_, _ = o.Raw("DELETE FROM notify_messages WHERE createTime < ?", cutoff).Exec()
}

func extractTitle(content string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		text := strings.TrimSpace(line)
		if text == "" {
			continue
		}
		text = strings.TrimLeft(text, "#> ")
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}
		return text
	}
	return "通知消息"
}

func detectLevel(moduleName string, content string) string {
	text := strings.ToLower(content)
	if strings.Contains(text, "fail") || strings.Contains(text, "error") ||
		strings.Contains(content, "失败") || strings.Contains(content, "错误") {
		return "error"
	}
	if strings.Contains(content, "报警") || strings.Contains(text, "warning") || strings.Contains(text, "notice") {
		return "warning"
	}
	if strings.Contains(text, "success") || strings.Contains(content, "成功") {
		return "info"
	}
	lowerModule := strings.ToLower(moduleName)
	if strings.Contains(lowerModule, "notice") || strings.Contains(lowerModule, "listen") || strings.Contains(lowerModule, "funding") {
		return "warning"
	}
	return "info"
}
