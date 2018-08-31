/**
 *
 *@author steve dingjc89@126.com
 *2018/8/16
 *@version V1.0
 *@copyright steve
 */
package libs

import (
	"quickstart/models"
	"strconv"
	"strings"
)

// 当前用户是否有权限
func Permission(userid int, pName string) bool {
	u := new(models.User)
	user, _ := u.GetUserById(userid)
	// 超级管理员有所有权限
	if user.Account == "admin" {
		return true
	}
	r := new(models.Role)
	roles, _ := r.GetRoleList(1,
		-1,
		map[string][]interface{}{"Id": []interface{}{"in", strings.Split(user.RoleIds, ",")}})
	var permission string
	for _, v := range roles {
		permission = v.PermissionIds
	}
	p := new(models.Permissions)
	pRow := p.GetOne(pName)

	if len(permission) > 0 {
		if strings.Contains(","+permission, ","+strconv.Itoa(pRow.Id)+",") {
			return true
		}
	}
	return false
}
