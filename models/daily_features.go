package models

type DailyFeature struct {
	ID                        int64   `orm:"column(id)" json:"id"`
	ScanDate                  string  `orm:"column(scan_date)" json:"scan_date"`
	Symbol                    string  `orm:"column(symbol)" json:"symbol"`
	MarkPrice                 float64 `orm:"column(mark_price)" json:"mark_price"`
	IndexPrice                float64 `orm:"column(index_price)" json:"index_price"`
	LastFundingRate           float64 `orm:"column(last_funding_rate)" json:"last_funding_rate"`
	NextFundingTime           int64   `orm:"column(next_funding_time)" json:"next_funding_time"`
	OpenInterest              float64 `orm:"column(open_interest)" json:"open_interest"`
	TopLongShortPositionRatio float64 `orm:"column(top_long_short_position_ratio)" json:"top_long_short_position_ratio"`
	TopLongShortAccountRatio  float64 `orm:"column(top_long_short_account_ratio)" json:"top_long_short_account_ratio"`
	GlobalLongShortRatio      float64 `orm:"column(global_long_short_ratio)" json:"global_long_short_ratio"`
	TakerBuySellRatio         float64 `orm:"column(taker_buy_sell_ratio)" json:"taker_buy_sell_ratio"`
	Basis                     float64 `orm:"column(basis)" json:"basis"`
	AnnualizedBasisRate       float64 `orm:"column(annualized_basis_rate)" json:"annualized_basis_rate"`
	FuturesPrice              float64 `orm:"column(futures_price)" json:"futures_price"`
	FundingRateInterval       int64   `orm:"column(funding_rate_interval)" json:"funding_rate_interval"`
	RawJson                   string  `orm:"column(raw_json);type(text)" json:"raw_json"`
	CreateTime                int64   `orm:"column(createTime)" json:"createTime"`
}

func (m *DailyFeature) TableName() string {
	return "daily_features"
}
