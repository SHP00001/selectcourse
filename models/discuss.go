package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	//"strconv"
	//	"errors"
	"time"
)

type Discuss struct {
	Id        int64     `json:"-"`
	Userid    int64     `orm:"index" json:"userid"`
	Courseid  int64     `orm:"index" json:"courseid"`
	Anonymous string    `orm:"size(100)",json:"anonymous"`
	Text      string    `orm:size(8000),json:"text"`
	Created   time.Time `orm:"auto_now_add;type(datetime)" json:"-"`
	Updated   time.Time `orm:"auto_now;type(datetime)" json:"-"`
}

func AddDiscuss(d Discuss) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(&d)
}
