package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"selectcourse/models"
)

type CourseStatusController struct {
	beego.Controller
}

func (this *CourseStatusController) Get() {
	var userid int64
	var courseid int64
	this.Ctx.Input.Bind(&userid, "userid")
	this.Ctx.Input.Bind(&courseid, "courseid")
	beego.Informational("userid: ", userid, "courseid: ", courseid, "ip", this.Ctx.Input.IP())
	status, err := models.GetCourseStatus(userid, courseid)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	remap := Map{"result": "ok"}
	remap["status"] = status
	re, err := json.Marshal(remap)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	this.Ctx.Output.Body(re)

	this.Ctx.Output.Body([]byte(`{"result":"ok","status":"subscribed"}`))
}

func (this *CourseStatusController) Put() {
	beego.Informational("login json: ", string(this.Ctx.Input.RequestBody), "ip", this.Ctx.Input.IP())
	var uc models.Usercourse
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &uc)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	err = models.UpdateCourseStatus(uc)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}

	this.Ctx.Output.Body([]byte(`{"result":"ok"}`))
}
