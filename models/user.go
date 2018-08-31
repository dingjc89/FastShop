package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int
	Account    string
	Pass       string
	RoleIds    string
	Name       string `orm:"null"`
	Age        int
	Phone      string `orm:"unique"`
	Emails     string `orm:"null"`
	Status     int    `orm:"default(1)"`
	CreatedAt  int64
	CreatedId  int
	ModifiedAt int64 `orm:"null"`
	ModifiedId int
	Deleted    int   `orm:"null;default(0)"`
	DeletedAt  int64 `orm:"null"`
}

func (u *User) TableName() string {
	return TableName("user")
}

func (u *User) TableIndex() [][]string {
	return [][]string{
		[]string{"Id", "Account"},
	}
}

func (u *User) GetUserByAccount(account string) (*User, error) {
	err := orm.NewOrm().QueryTable(TableName("user")).Filter("account", account).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) GetUserById(id int) (*User, error) {
	err := orm.NewOrm().QueryTable(TableName("user")).Filter("id", id).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) GetUserList(page, limit int, filter map[string][]interface{}) ([]*User, int64) {
	offset := (page - 1) * limit
	_filter := Filter(filter)
	query := orm.NewOrm().QueryTable(TableName("user"))
	for key, value := range _filter {
		query = query.Filter(key, value...)
	}
	uList := new([]*User)
	query.Limit(limit, offset).All(uList)
	total, _ := query.Count()
	return *uList, total
}

func (u *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) Insert() error {
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}
