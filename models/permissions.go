/**
 *
 *@author steve dingjc89@126.com
 *2018/8/16
 *@version V1.0
 *@copyright steve
 */
package models

import "github.com/astaxie/beego/orm"

type Permissions struct {
	Id         int
	Name       string
	Pn         string
	Pid        int
	Img        string
	Order      int
	IsShow     int
	CreatedId  int
	Url        string `orm:"null"`
	ModifiedId int    `orm:"null"`
	CreatedAt  int64
	ModifiedAt int64 `orm:"null"`
}

func (p *Permissions) TableName() string {
	return TableName("permissions")
}

func (p *Permissions) GetPermissionList(filter map[string][]interface{}) []*Permissions {
	query := orm.NewOrm().QueryTable(TableName("permissions"))
	_filter := Filter(filter)
	for key, value := range _filter {
		query = query.Filter(key, value...)
	}
	list := make([]*Permissions, 0)
	query.OrderBy("pid", "order").All(&list)
	return list
}

func (p *Permissions) GetOne(pName string) *Permissions {
	orm.NewOrm().QueryTable(TableName("permissions")).Filter("Pn", pName).One(p)
	return p
}
