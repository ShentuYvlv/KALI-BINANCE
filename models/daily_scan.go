package models

type DailyScan struct {
	ID                 int64   `orm:"column(id)" json:"id"`
	ScanDate           string  `orm:"column(scan_date)" json:"scan_date"`
	Symbol             string  `orm:"column(symbol)" json:"symbol"`
	Direction          string  `orm:"column(direction)" json:"direction"` // gain/lose
	Rank               int64   `orm:"column(rank)" json:"rank"`
	PriceChangePercent float64 `orm:"column(price_change_percent)" json:"price_change_percent"`
	QuoteVolume        float64 `orm:"column(quote_volume)" json:"quote_volume"`
	OpenTime           int64   `orm:"column(open_time)" json:"open_time"`
	CloseTime          int64   `orm:"column(close_time)" json:"close_time"`
	CreateTime         int64   `orm:"column(createTime)" json:"createTime"`
}

func (m *DailyScan) TableName() string {
	return "daily_scan"
}
