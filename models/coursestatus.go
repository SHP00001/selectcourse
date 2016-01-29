package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	//"strconv"
	"errors"
	"time"
)

type Usercourse struct {
	Id       int64     `json:"-"`
	Userid   int64     `orm:"index" json:"userid"`
	Courseid int64     `orm:"index" json:"courseid"`
	Status   string    `orm:"size(100);index" json:"status"`
	Adminid  int64     `orm:"-" json:"adminid"`
	Created  time.Time `orm:"auto_now_add;type(datetime)" json:"-"`
	Updated  time.Time `orm:"auto_now;type(datetime)" json:"-"`
}

func UpdateCourseSelected(cid int64, count int) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("course").Filter("id", cid).Update(orm.Params{
		"selected": orm.ColValue(orm.ColAdd, count)})
	return err
}

func GetUserCourse(uid int64, cid int64) (Usercourse, error) {
	o := orm.NewOrm()
	var uc Usercourse
	err := o.QueryTable("usercourse").Filter("userid", uid).Filter("courseid", cid).One(&uc)
	return uc, err
}

func GetUserCourseByUserId(uid int64) ([]Usercourse, error) {
	o := orm.NewOrm()
	var uc []Usercourse
	_, err := o.QueryTable("usercourse").Filter("userid", uid).All(&uc)
	return uc, err
}

func GetCourseStatus(uid int64, cid int64) (string, error) {
	uc, err := GetUserCourse(uid, cid)
	if err != nil {
		return "", err
	}
	if uc.Status == "subscribed" || uc.Status == "approval" {
		return uc.Status, nil
	}
	c, err := GetCourseByID(cid)
	if err != nil {
		return "", err
	}
	if c.Total >= c.Selected {
		return "available", nil
	} else {
		return "unavailable", nil
	}
}

func UpdateCourseStatus(uc Usercourse) error {
	o := orm.NewOrm()
	couse, err := GetCourseByID(uc.Courseid)
	if err != nil {
		return err
	}
	user, err := GetUserByID(uc.Userid)
	if err != nil {
		return err
	}
	if uc.Status == "del" {
		if CheckRole(uc.Adminid, "admin") {
			_, err := o.QueryTable("course").Filter("id", uc.Courseid).Delete()
			return err
		} else {
			return errors.New("需要管理员权限")
		}
	} else if CheckRole(uc.Adminid, "admin") {
		if uc.Status == "available" {
			err = UpdateCourseSelected(uc.Courseid, -1)
			if err != nil {
				return err
			}
			_, err = o.QueryTable("usercourse").Filter("userid", uc.Userid).Filter("courseid", uc.Courseid).Delete()
			if err != nil {
				return err
			}
			_, err = o.QueryTable("usercourse").Filter("userid", uc.Userid).Filter("courseid", uc.Courseid).Update(orm.Params{
				"status": uc.Status, "updated": time.Now()})
			if err != nil {
				return err
			}
			return nil
		}
		if uc.Status == "subscribed" {
			if user.Coursehour <= 0 {
				return errors.New("课时不足")
			}
			err := UpdateUserCourseHour(user.Id, -1)
			if err != nil {
				return err
			}
		}
		_, err = o.QueryTable("usercourse").Filter("userid", uc.Userid).Filter("courseid", uc.Courseid).Update(orm.Params{
			"status": uc.Status, "updated": time.Now()})
		if err != nil {
			return err
		}
		return nil
	} else if couse.Type == "vip" {
		if !CheckRole(uc.Userid, "student") {
			return errors.New("需要学生身份才可选择vip课程")
		} else if uc.Status == "subscribed" {
			return errors.New("vip课程需要先进入待审批状态")
		}
	} else if couse.Type != "vip" && uc.Status == "approval" {
		return errors.New("非vip课程不存在审批过程")
	}
	exist := o.QueryTable("usercourse").Filter("userid", uc.Userid).Filter("courseid", uc.Courseid).Exist()
	//fmt.Println(exist)
	if exist {
		return errors.New("禁止重复提交")
	}
	err = UpdateCourseSelected(uc.Courseid, 1)
	if err != nil {
		return err
	}
	_, err = AddUserCourse(uc)
	return err

}

func AddUserCourse(uc Usercourse) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(&uc)
}
