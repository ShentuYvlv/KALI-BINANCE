package models

type FeatureTimeseries struct {
	ID         int64   `orm:"column(id)" json:"id"`
	Symbol     string  `orm:"column(symbol)" json:"symbol"`
	Metric     string  `orm:"column(metric)" json:"metric"`
	Timestamp  int64   `orm:"column(timestamp)" json:"timestamp"`
	Value      float64 `orm:"column(value)" json:"value"`
	ExtraJson  string  `orm:"column(extra_json);type(text)" json:"extra_json"`
	CreateTime int64   `orm:"column(createTime)" json:"createTime"`
}

func (m *FeatureTimeseries) TableName() string {
	return "feature_timeseries"
}
