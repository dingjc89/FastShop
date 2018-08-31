/**
 *
 *@author steve dingjc89@126.com
 *2018/8/16
 *@version V1.0
 *@copyright steve
 */
package models

import "github.com/astaxie/beego/orm"

type Logs struct {
	Id        int
	Ip        string `orm:"size(20)"`
	Actions   string
	Detail    string `orm:"type(text)"`
	CreatedId int
	CreatedAt int64
}

func (l *Logs) TableName() string {
	return TableName("logs")
}

func AddLog(l *Logs) (int64, error) {
	return orm.NewOrm().Insert(l)
}
