package models

import (
	//	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	//"strconv"
	"errors"
	"time"
)

//var (
//UserList map[string]*User
//)

type User struct {
	Id         int64     `json:"userid"`
	Username   string    `orm:"size(100);index;unique" json:"username"`
	Password   string    `orm:"size(100)" json:"password"`
	Name       string    `orm:"size(100)" json:"name"`
	Phone      int       `orm:"size(100)" json:"phone"`
	Age        int       `orm:"size(100)" json:"age"`
	Role       string    `orm:"size(100)" json:"role"`
	Coursehour int       `orm:"size(100)" json:"-"`
	Created    time.Time `orm:"auto_now_add;type(datetime)" json:"-"`
	Updated    time.Time `orm:"auto_now;type(datetime)" json:"-"`
	Result     string    `orm:"-" json:"result"`
}

func UpdateUserCourseHour(uid int64, count int) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("user").Filter("id", uid).Update(orm.Params{
		"coursehour": orm.ColValue(orm.ColAdd, count)})
	return err
}

func AddUser(u User) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(&u)
}

func GetUserByUserName(uname string) (User, error) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("username", uname).One(&user)
	return user, err

}

func GetUserByID(uid int64) (User, error) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("id", uid).One(&user)
	return user, err

}

func CheckRole(uid int64, role string) bool {
	u, _ := GetUserByID(uid)
	return u.Role == role
}

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
		"password": np,"updated":time.Now()})
	return nil
}
