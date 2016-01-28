package controllers

import (
	"github.com/astaxie/beego"
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
