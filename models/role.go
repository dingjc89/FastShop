/**
 *
 *@author steve dingjc89@126.com
 *2018/8/16
 *@version V1.0
 *@copyright steve
 */
package models

import "github.com/astaxie/beego/orm"

type Role struct {
	Id            int
	PermissionIds string
	Name          string
	Detail        string `orm:"type(text)"`
	CreatedId     int
	ModifiedId    int `orm:"null;default(0)"`
	CreatedAt     int64
	ModifiedAt    int64 `orm:"null"`
	Status        int   `orm:"default(0)"`
	DeletedAt     int64 `orm:"null"`
	DeletedId     int   `orm:"null;default(0)"`
}

func (r *Role) TableName() string {
	return TableName("role")
}

func (r *Role) GetRoleList(page, limit int, filter map[string][]interface{}) ([]*Role, int64) {
	filter["status"] = []interface{}{0} // 默认查询条件

	offset := (page - 1) * limit
	_filter := Filter(filter)
	query := orm.NewOrm().QueryTable(TableName("role"))
	for key, value := range _filter {
		query = query.Filter(key, value...)
	}
	role := new([]*Role)
	query.Limit(limit, offset).All(role)
	total, _ := query.Count()
	return *role, total
}

func (r *Role) Update(fields ...string) error {
	_, err := orm.NewOrm().Update(r, fields...)
	if err != nil {
		return err
	}
	return nil
}

func (r *Role) Add() error {
	_, err := orm.NewOrm().Insert(r)
	if err != nil {
		return err
	}
	return nil
}

func (r *Role) GetRoleById(id int) (*Role, error) {
	err := orm.NewOrm().QueryTable(TableName("role")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
