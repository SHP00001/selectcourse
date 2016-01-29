package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"selectcourse/models"
)

type DiscussController struct {
	beego.Controller
}

func (this *DiscussController) Get() {
	page := this.Input().Get("page")
	count := this.Input().Get("count")
	courseid := this.Input().Get("courseid")
	beego.Informational("page: ", page, "count: ", count, "courseid: ", courseid, "ip", this.Ctx.Input.IP())
	this.Ctx.Output.Body([]byte(`{"result":"ok","discuss":[{"name":"**","text":"**","ctime":21321321},{"name":"**","text":"**","ctime":21321321}]}`))
}

func (this *DiscussController) Post() {
	beego.Informational("post json: ", string(this.Ctx.Input.RequestBody), "ip", this.Ctx.Input.IP())
	var discuss models.Discuss
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &discuss)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	_, err = models.AddDiscuss(discuss)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	this.Ctx.Output.Body([]byte(`{"result":"ok"}`))
}

func (this *DiscussController) MsgGet() {
	page := this.Input().Get("page")
	count := this.Input().Get("count")
	courseid := this.Input().Get("courseid")
	beego.Informational("page: ", page, "count: ", count, "courseid: ", courseid, "ip", this.Ctx.Input.IP())
	this.Ctx.Output.Body([]byte(`{"result":"ok","discuss":[{"name":"**","text":"**","ctime":21321321},{"name":"**","text":"**","ctime":21321321}]}`))
}

func (this *DiscussController) MsgPost() {
	beego.Informational("msgpost json: ", string(this.Ctx.Input.RequestBody), "ip", this.Ctx.Input.IP())
	this.Ctx.Output.Body([]byte(`{"result":"ok"}`))
}
