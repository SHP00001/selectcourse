package main

import (
	_ "selectcourse/docs"
	_ "selectcourse/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"selectcourse/controllers"
	"selectcourse/models"
)

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@/selectcourse?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(models.User),new(models.Course),new(models.Usercourse))

	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Router("/api/login", &controllers.LoginController{})
	beego.Router("/api/register", &controllers.AccountController{})
	beego.Router("/api/password", &controllers.AccountController{})
	beego.Router("/api/ad", &controllers.ADController{})
	beego.Router("/api/courses", &controllers.CourseController{}, "get:GetByTime")
	beego.Router("/api/mycourses", &controllers.CourseController{}, "get:GetByUserID")
	beego.Router("/api/newcourse", &controllers.CourseController{}, "post:NewCourse")
	beego.Router("/api/newcourse", &controllers.CourseController{}, "post:NewCourse")
	beego.Router("/api/newcourse", &controllers.CourseController{}, "post:NewCourse")
	beego.Router("/api/course/status", &controllers.CourseStatusController{})
	beego.Router("/api/course/discuss", &controllers.DiscussController{})
	beego.Router("/api/discuss", &controllers.DiscussController{}, "get:MsgGet;post:MsgPost")
	beego.Router("/api/admin/courses", &controllers.CourseController{}, "get:AdminGet")
	beego.Router("/api/admin/courseshour", &controllers.AccountController{}, "post:AddCourseHour")
	beego.Run()
}
