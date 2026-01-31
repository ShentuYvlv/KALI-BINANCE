package models

type BacktestT1 struct {
	ID         int64   `orm:"column(id)" json:"id"`
	ScanDate   string  `orm:"column(scan_date)" json:"scan_date"`
	Symbol     string  `orm:"column(symbol)" json:"symbol"`
	Direction  string  `orm:"column(direction)" json:"direction"` // gain/lose
	Rank       int64   `orm:"column(rank)" json:"rank"`
	T1Open     float64 `orm:"column(t1_open)" json:"t1_open"`
	T1Close    float64 `orm:"column(t1_close)" json:"t1_close"`
	T1Return   float64 `orm:"column(t1_return)" json:"t1_return"`
	CreateTime int64   `orm:"column(createTime)" json:"createTime"`
}

func (m *BacktestT1) TableName() string {
	return "backtest_t1"
}
