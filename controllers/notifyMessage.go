package controllers

import (
	"strconv"

	"go_binance_futures/models"
	"go_binance_futures/utils"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type NotifyMessageController struct {
	web.Controller
}

func (ctrl *NotifyMessageController) Get() {
	pageStr := ctrl.GetString("page", "1")
	limitStr := ctrl.GetString("limit", "50")
	module := ctrl.GetString("module")
	channel := ctrl.GetString("channel")
	level := ctrl.GetString("level")
	keyword := ctrl.GetString("keyword")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 50
	}
	offset := (page - 1) * limit

	o := orm.NewOrm()
	query := o.QueryTable("notify_messages")
	countQuery := o.QueryTable("notify_messages")
	if module != "" {
		query = query.Filter("module", module)
		countQuery = countQuery.Filter("module", module)
	}
	if channel != "" {
		query = query.Filter("channel", channel)
		countQuery = countQuery.Filter("channel", channel)
	}
	if level != "" {
		query = query.Filter("level", level)
		countQuery = countQuery.Filter("level", level)
	}
	if keyword != "" {
		query = query.Filter("content__contains", keyword)
		countQuery = countQuery.Filter("content__contains", keyword)
	}

	var list []models.NotifyMessage
	_, err := query.OrderBy("-id").Limit(limit, offset).All(&list)
	if err != nil {
		ctrl.Ctx.Resp(utils.ResJson(400, nil, err.Error()))
		return
	}
	total, err := countQuery.Count()
	if err != nil {
		ctrl.Ctx.Resp(utils.ResJson(400, nil, err.Error()))
		return
	}

	ctrl.Ctx.Resp(map[string]interface{}{
		"code": 200,
		"data": map[string]interface{}{
			"total": total,
			"list":  list,
		},
		"msg": "success",
	})
}

func (ctrl *NotifyMessageController) Modules() {
	o := orm.NewOrm()
	var rows []orm.Params
	_, err := o.Raw("SELECT module, COUNT(*) as count FROM notify_messages GROUP BY module ORDER BY count DESC").Values(&rows)
	if err != nil {
		ctrl.Ctx.Resp(utils.ResJson(400, nil, err.Error()))
		return
	}
	ctrl.Ctx.Resp(map[string]interface{}{
		"code": 200,
		"data": rows,
		"msg": "success",
	})
}

func (ctrl *NotifyMessageController) Delete() {
	id := ctrl.Ctx.Input.Param(":id")
	msg := new(models.NotifyMessage)
	intId, _ := strconv.ParseInt(id, 10, 64)
	msg.ID = intId
	o := orm.NewOrm()
	_, err := o.Delete(msg)
	if err != nil {
		ctrl.Ctx.Resp(utils.ResJson(400, nil, "删除错误"))
		return
	}
	ctrl.Ctx.Resp(utils.ResJson(200, nil))
}

func (ctrl *NotifyMessageController) Clear() {
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM notify_messages").Exec()
	if err != nil {
		ctrl.Ctx.Resp(utils.ResJson(400, nil, "清空错误"))
		return
	}
	ctrl.Ctx.Resp(utils.ResJson(200, nil))
}
