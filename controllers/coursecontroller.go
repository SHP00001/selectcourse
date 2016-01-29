package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"selectcourse/models"
)

type CourseController struct {
	beego.Controller
}

func (this *CourseController) GetByTime() {
	var year int
	var month int
	this.Ctx.Input.Bind(&year, "year")
	this.Ctx.Input.Bind(&month, "month")
	beego.Informational("year:", year, " month:", month)
	cl, err := models.GetCourseByTime(year, month)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	remap := Map{"result": "ok"}
	remap["course"] = cl
	re, err := json.Marshal(remap)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	this.Ctx.Output.Body(re)
	//this.Ctx.Output.Body([]byte(`{"result":"ok","courses":[{"courseid":123,"coursename":"**","teachername":"**","year":2016,"month":1,"day":20,"start":111111,"end":11111,"details":"**","coursetype":"vip"},{"courseid":123,"coursename":"**","teachername":"**","year":2016,"month":1,"day":20,"start":111111,"end":11111,"details":"**","coursetype":"vip"}]}`))
}

func (this *CourseController) GetByUserID() {
	var userid int64
	this.Ctx.Input.Bind(&userid, "userid")
	beego.Informational("userid:", userid)
	user, err := models.GetUserByID(userid)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	cl, err := models.GetCourseByUserID(userid)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	remap := Map{"result": "ok"}
	remap["course"] = cl
	remap["coursehour"] = user.Coursehour
	re, err := json.Marshal(remap)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	this.Ctx.Output.Body(re)
}

func (this *CourseController) NewCourse() {
	beego.Informational("newcourse:", string(this.Ctx.Input.RequestBody), "ip", this.Ctx.Input.IP())
	var course models.Course
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &course)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	beego.Informational("course:", course)
	if !models.CheckRole(course.Userid, "admin") {
		ErrMap["msg"] = "权限不足"
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	//c:= models.Course{}
	cid, err := models.AddCourse(course)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	remap := Map{"result": "ok"}
	remap["courseid"] = cid
	remap["coursename"] = course.Coursename
	re, err := json.Marshal(remap)
	if err != nil {
		beego.Error("err:", err.Error())
		ErrMap["msg"] = err.Error()
		re, _ := json.Marshal(ErrMap)
		this.Ctx.Output.Body(re)
		return
	}
	this.Ctx.Output.Body(re)
	//this.Ctx.Output.Body([]byte(`{"result":"ok","courseid":"123","coursename":"abc"}`))
}

func (this *CourseController) AdminGet() {
	userid := this.Input().Get("userid")
	beego.Informational("userid:", userid)
	this.Ctx.Output.Body([]byte(`{"result":"ok","courses":[{"courseid":123,"coursename":"**","teachername":"**","year":2016,"month":1,"day":20,"start":111111,"end":11111,"details":"**","coursetype":"vip","studentname":"123","studentphone":12345,"parentsname":"1234","parentsphone":1234},{"courseid":123,"coursename":"**","teachername":"**","year":2016,"month":1,"day":20,"start":111111,"end":11111,"details":"**","coursetype":"vip","studentname":"123","studentphone":12345,"parentsname":"1234","parentsphone":1234}]}`))
}
