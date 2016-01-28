package controllers

import (
	"github.com/astaxie/beego"
)

type ADController struct {
	beego.Controller
}

func (this *ADController) Get() {
	this.Ctx.Output.Body([]byte(`{"result":"ok","url":"***"}`))
}
