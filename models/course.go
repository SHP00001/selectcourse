package models

import (
	//	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	//"strconv"
	//	"errors"
	"time"
)

//var (
//UserList map[string]*User
//)

type Course struct {
	Id          int64     `json:"courseid"`
	Coursename  string    `orm:"size(100)" json:"coursename"`
	Teachername string    `orm:"size(100)" json:"teachername"`
	Year        int       `orm:"index" json:"year"`
	Month       int       `orm:"index" json:"month"`
	Day         int       `orm:"index" json:"day"`
	Stat        int       `orm:"size(100);index" json:"start"`
	End         int       `orm:"size(100)" json:"end"`
	Details     string    `orm:"size(100)" json:"details"`
	Type        string    `orm:"size(100)" json:"coursetype"`
	Total       int       `json:"total"`
	Selected    int       `json:"select"`
	Created     time.Time `orm:"auto_now_add;type(datetime)" json:"-"`
	Updated     time.Time `orm:"auto_now;type(datetime)" json:"-"`
	Userid      int64     `json:"userid"`
}

func GetCourseByUserID(uid int64)([]Course,error){
	var re []Course
	uclist,err:=GetUserCourseByUserId(uid)
	if err!=nil{
		return nil,err
	}
	for _,v:= range uclist{
		c,err:=GetCourseByID(v.Courseid)
		if err!=nil{
			return nil,err
		}
		re=append(re,c)
	}
	return re,nil

}

func AddCourse(c Course) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(&c)
}

func GetCourseByTime(year int, month int) ([]Course, error) {
	o := orm.NewOrm()
	var listC []Course
	_, err := o.QueryTable("course").Filter("year", year).Filter("month", month).All(&listC)
	if err != nil && err.Error() == "<QuerySeter> no row found" {
		return listC, nil
	}
	return listC, err

}

func GetCourseByID(cid int64) (Course, error) {
	o := orm.NewOrm()
	var c Course
	err := o.QueryTable("course").Filter("id", cid).One(&c)
	return c, err

}

/*
func UpdatePassword(uid int64, op string, np string) error {
	o := orm.NewOrm()
	u, err := GetUserByID(uid)
	if err != nil {
		return err
	}
	if u.Password != op {
		return errors.New("密码错误")
	}
	_, err = o.QueryTable("user").Filter("id", uid).Update(orm.Params{
		"password": np})
	return nil
}*/
