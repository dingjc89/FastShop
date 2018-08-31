/**
 *
 *@author steve dingjc89@126.com
 *2018/8/17
 *@version V1.0
 *@copyright steve
 */
package controllers

import (
	"quickstart/models"
	"time"

	"strings"

	"strconv"

	"github.com/astaxie/beego/logs"
)

type RoleController struct {
	BaseController
}

func (self *RoleController) List() {
	self.display("role/list", true)
}

func (self *RoleController) Table() {
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 10
	}
	filter := make(map[string][]interface{})
	name := strings.TrimSpace(self.GetString("name"))
	if len(name) > 0 {
		filter["name"] = []interface{}{"icontains", name}
	}
	role := new(models.Role)
	roles, count := role.GetRoleList(page, limit, filter)

	data := make([]map[string]interface{}, len(roles))
	for k, v := range roles {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["name"] = v.Name
		row["memo"] = v.Detail
		data[k] = row
	}
	self.AjaxList(SUCC, "", count, data)

}

func (self *RoleController) Edit() {
	id, _ := self.GetInt("id")
	roleM := new(models.Role)
	role, err := roleM.GetRoleById(id)
	if err != nil {
		logs.Error(id, "角色ID："+string(id)+"没有找到对应的数据")
	}
	data := make(map[string]interface{})
	data["id"] = role.Id
	data["name"] = role.Name
	data["detail"] = role.Detail
	// 角色权限
	permissionSlice := make([]int, 0)
	rolep := strings.Split(role.PermissionIds, ",")
	for _, v := range rolep {
		p, _ := strconv.Atoi(v)
		permissionSlice = append(permissionSlice, p)
	}

	self.Data["role"] = data
	self.Data["permission"] = getPermission()
	self.Data["checked"] = permissionSlice
	self.Data["zTree"] = true
	self.display("role/edit", true)

}

func getPermission() []map[string]interface{} {
	// 全部权限
	permission := new(models.Permissions)
	permissions := permission.GetPermissionList(map[string][]interface{}{})
	pdata := make([]map[string]interface{}, len(permissions))
	for k, v := range permissions {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["name"] = v.Name
		row["pId"] = v.Pid
		pdata[k] = row
	}
	return pdata
}

func (self *RoleController) AjaxDel() {
	id, _ := self.GetInt("id")
	role := new(models.Role)
	role.Id = id
	role.Status = 1
	role.DeletedAt = time.Now().Unix()
	role.DeletedId = self.userId
	if err := role.Update("Status", "DeletedAt", "DeletedId"); err != nil {
		self.AjaxJson(FAIL, "删除失败")
	}
	self.AjaxJson(SUCC, "删除成功")
}

func (self *RoleController) AjaxSave() {

	id, _ := self.GetInt("id")
	fields := make([]string, 0)
	name := self.GetString("name")
	detail := self.GetString("detail")
	checked := self.GetString("checked") // 选中的权限

	role := new(models.Role)
	role.Name = name
	role.Detail = detail
	role.PermissionIds = checked
	if id > 0 {
		role.Id = id
		role.ModifiedId = self.userId
		role.ModifiedAt = time.Now().Unix()
		fields = []string{"Id", "Name", "Detail", "PermissionIds", "ModifiedId", "ModifiedAt"}
		if err := role.Update(fields...); err != nil {
			self.AjaxJson(FAIL, "保存失败")
		}
	} else {
		role.CreatedAt = time.Now().Unix()
		role.CreatedId = self.userId
		if err := role.Add(); err != nil {
			self.AjaxJson(FAIL, "保存失败")
		}
	}

	self.AjaxJson(SUCC, "保持成功")
}

func (self *RoleController) Add() {
	self.Data["zTree"] = true

	self.Data["permissions"] = getPermission()
	self.display("role/add", true)
}
