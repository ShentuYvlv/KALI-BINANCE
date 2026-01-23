package models

import (
	"github.com/beego/beego/v2/client/orm"
)

func EnsureListenSchema() {
	o := orm.NewOrm()
	// best-effort add columns for merged notice fields
	_ = addColumnIgnoreErr(o, "listen_symbols", "notice_price", "TEXT")
	_ = addColumnIgnoreErr(o, "listen_symbols", "has_notice", "INTEGER DEFAULT 0")
	_ = addColumnIgnoreErr(o, "listen_symbols", "auto_order", "INTEGER DEFAULT 0")
	_ = addColumnIgnoreErr(o, "listen_symbols", "profit_price", "TEXT")
	_ = addColumnIgnoreErr(o, "listen_symbols", "loss_price", "TEXT")
	_ = addColumnIgnoreErr(o, "listen_symbols", "leverage", "INTEGER DEFAULT 0")
	_ = addColumnIgnoreErr(o, "listen_symbols", "marginType", "TEXT")
	_ = addColumnIgnoreErr(o, "listen_symbols", "tickSize", "TEXT")
	_ = addColumnIgnoreErr(o, "listen_symbols", "stepSize", "TEXT")
	_ = addColumnIgnoreErr(o, "listen_symbols", "usdt", "TEXT")
	_ = addColumnIgnoreErr(o, "listen_symbols", "side", "TEXT")
	_ = addColumnIgnoreErr(o, "listen_symbols", "quantity", "TEXT")
}

func MigrateNoticeToListen() {
	o := orm.NewOrm()
	var notices []NoticeSymbols
	_, err := o.QueryTable("notice_symbols").All(&notices)
	if err != nil || len(notices) == 0 {
		return
	}

	for _, notice := range notices {
		exists := o.QueryTable("listen_symbols").
			Filter("symbol", notice.Symbol).
			Filter("type", notice.Type).
			Filter("listen_type", "price_notice").
			Exist()
		if exists {
			continue
		}
		listen := ListenSymbols{
			Symbol: notice.Symbol,
			Enable: notice.Enable,
			Type: notice.Type,
			ListenType: "price_notice",
			NoticePrice: notice.NoticePrice,
			HasNotice: notice.HasNotice,
			AutoOrder: notice.AutoOrder,
			ProfitPrice: notice.ProfitPrice,
			LossPrice: notice.LossPrice,
			Leverage: notice.Leverage,
			MarginType: notice.MarginType,
			TickSize: notice.TickSize,
			StepSize: notice.StepSize,
			Usdt: notice.Usdt,
			Side: notice.Side,
			Quantity: notice.Quantity,
			NoticeLimitMin: 5,
			LastNoticeTime: 0,
			LastNoticeType: "",
		}
		_, _ = o.Insert(&listen)
	}
}

func addColumnIgnoreErr(o orm.Ormer, table string, column string, columnType string) error {
	_, err := o.Raw("ALTER TABLE " + table + " ADD COLUMN " + column + " " + columnType).Exec()
	return err
}
