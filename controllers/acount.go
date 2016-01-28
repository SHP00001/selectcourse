package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"selectcourse/models"
	"strings"
)

type Pwd struct {
	Id int64  `json:"userid"`
	Np string `json:"newpassword"`
	Op string `json:"oldpassword"`
}
type Map map[string]interface{}

var ErrMap = Map{"result": "error", "msg": ""}

//var ReMap = Map{"result": "ok"}

type AccountController struct {
	beego.Controller
}

func (this *AccountController) Get() {
	this.Ctx.Output.Body([]byte("account get"))
}

func (this *AccountController) AddCourseHour() {
	beego.Informational("login json: ", string(this.Ctx.Input.RequestBody), "ip:", this.Ctx.Input.IP())
	type ach struct {
		Uname      string `json:"studentusername"`
		Coursehour int    `json:"coursehour"`
	}
	var ac ach
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &ac)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	user, err := models.GetUserByUserName(ac.Uname)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	err = models.UpdateUserCourseHour(user.Id, ac.Coursehour)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	this.Ctx.Output.Body([]byte(`{"result":"ok"}`))
}

func (this *AccountController) Post() {
	beego.Informational("login json: ", string(this.Ctx.Input.RequestBody), "ip:", this.Ctx.Input.IP())
	var user models.User
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	//this.Ctx.Output.Body([]byte(`{"result":"ok","userid":123,"role":"student","name":"lu"}`))
	id, err := models.AddUser(user)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		if strings.Contains(err.Error(), "Duplicate entry") {
			ErrMap["msg"] = "该用户名已被注册"
		}
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	user.Id = id
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

func (this *AccountController) Put() {
	beego.Informational("login json: ", string(this.Ctx.Input.RequestBody), "ip:", this.Ctx.Input.IP())
	var pwd Pwd
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &pwd)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	err = models.UpdatePassword(pwd.Id, pwd.Op, pwd.Np)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	this.Ctx.Output.Body([]byte(`{"result":"ok"}`))
}
