/**
 *
 *@author steve dingjc89@126.com
 *2018/8/17
 *@version V1.0
 *@copyright steve
 */
package controllers

import (
	"quickstart/libs"
	"quickstart/models"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
	BaseController
}

func (self *UserController) List() {
	self.Data["pageTitle"] = "用户管理"
	self.display("user/list", true)
}

func (self *UserController) Table() {
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 10
	}
	name := strings.TrimSpace(self.GetString("name"))
	filter := make(map[string][]interface{}, 0)
	if len(name) > 0 {
		filter["name"] = []interface{}{"icontains", name}
	}
	user := new(models.User)
	list, total := user.GetUserList(page, limit, filter)
	data := make([]map[string]interface{}, len(list))
	for k, v := range list {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["account"] = v.Account
		row["real_name"] = v.Name
		row["phone"] = v.Phone
		row["email"] = v.Emails
		row["age"] = v.Age
		row["status"] = v.Status
		data[k] = row
	}
	self.AjaxList(SUCC, "", total, data)
}

func (self *UserController) AjaxDel() {
	id, _ := self.GetInt("id")
	statusStr := self.GetString("status")
	status := 0
	if statusStr == "enable" {
		status = 1
	}
	user := new(models.User)
	user.Id = id
	user.Status = status
	if err := user.Update("status"); err != nil {
		self.AjaxJson(FAIL, "失败")
	}
	self.AjaxJson(SUCC, "成功")

}

func (self *UserController) Edit() {
	id, _ := self.GetInt("id")
	userM := new(models.User)
	user, _ := userM.GetUserById(id)
	userMap := make(map[string]interface{})
	userMap["id"] = user.Id
	userMap["account"] = user.Account
	userMap["name"] = user.Name
	userMap["phone"] = user.Phone
	userMap["email"] = user.Emails
	userMap["age"] = user.Age

	checked := "," + user.RoleIds + ","

	r := new(models.Role)
	roles, _ := r.GetRoleList(1, -1, map[string][]interface{}{})
	list := make([]map[string]interface{}, len(roles))
	for k, v := range roles {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["name"] = v.Name
		row["checked"] = 0
		if strings.Contains(checked, ","+strconv.Itoa(v.Id)+",") {
			row["checked"] = 1
		}
		list[k] = row
	}
	self.Data["role"] = list

	self.Data["user"] = userMap
	self.display("user/edit", true)
}

func (self *UserController) AjaxSave() {
	user := new(models.User)
	user.Name = strings.TrimSpace(self.GetString("name"))
	user.Emails = strings.TrimSpace(self.GetString("email"))
	user.Age, _ = self.GetInt("age")
	user.Phone = strings.TrimSpace(self.GetString("phone"))
	user.RoleIds = strings.TrimSpace(self.GetString("roleids"))

	fields := []string{"Id", "Name", "Emails", "Age", "Phone", "ModifiedAt", "RoleIds"}
	user.Id, _ = self.GetInt("userid")

	if user.Id > 0 { // 编辑
		user.ModifiedAt = time.Now().Unix()
		user.ModifiedId = self.userId
		if i, _ := self.GetInt("reset_pwd"); i == 1 {
			user.Pass = libs.Password("admin123")
			fields = append(fields, "Pass")
			if err := user.Update(fields...); err != nil {
				self.AjaxJson(FAIL, "保存失败")
			}
		} else {
			if err := user.Update(fields...); err != nil {
				self.AjaxJson(FAIL, "保存失败")
			}
		}
	} else {
		user.Pass = libs.Password(strings.TrimSpace(self.GetString("password")))
		user.CreatedAt = time.Now().Unix()
		user.CreatedId = self.userId
		user.Account = strings.TrimSpace(self.GetString("account"))
		if err := user.Insert(); err != nil {
			self.AjaxJson(FAIL, "保存失败")
		}
	}

	self.AjaxJson(SUCC, "保存成功")
}

func (self *UserController) Add() {
	r := new(models.Role)
	roles, _ := r.GetRoleList(1, -1, map[string][]interface{}{})
	list := make([]map[string]interface{}, len(roles))
	for k, v := range roles {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["name"] = v.Name
		list[k] = row
	}
	self.Data["role"] = list
	self.display("user/add", true)
}
