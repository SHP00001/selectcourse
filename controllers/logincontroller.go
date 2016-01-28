package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"selectcourse/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.Ctx.Output.Body([]byte("123456"))
}

func (this *LoginController) Post() {
	beego.Informational("login json: ", string(this.Ctx.Input.RequestBody), "ip", this.Ctx.Input.IP())
	var uandp models.User
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &uandp)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	user, err := models.GetUserByUserName(uandp.Username)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = "用户名或密码错误"
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	if user.Password != uandp.Password {
		ErrMap["msg"] = "用户名或密码错误"
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	user.Result = "ok"
	re, err := json.Marshal(user)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	this.Ctx.Output.Body(re)
}
