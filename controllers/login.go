/**
 *
 *@author steve dingjc89@126.com
 *2018/8/15
 *@version V1.0
 *@copyright steve
 */
package controllers

import (
	"quickstart/libs"
	"quickstart/models"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type LoginController struct {
	BaseController
}

func (self *LoginController) Login() {
	self.display("login/login")
}

func (self *LoginController) LoginIn() {
	if self.isPost() {
		account := strings.TrimSpace(self.GetString("username"))
		pass := strings.TrimSpace(self.GetString("password"))
		user := new(models.User)
		if account != "" && pass != "" {
			user, err := user.GetUserByAccount(account)
			if err != nil || user.Pass != libs.Password(pass) {
				self.AjaxJson(FAIL, "帐号或密码错误")
			} else if user.Status != 1 {
				self.AjaxJson(FAIL, "账号已禁用")
			}
			l := new(models.Logs)
			l.Ip = self.getClientIP()
			l.Actions = "用户登录"
			l.Detail = "系统"
			l.CreatedId = user.Id
			l.CreatedAt = time.Now().Unix()
			models.AddLog(l)
			self.SetSession("account", account)
			self.SetSession("userid", user.Id)
			self.AjaxJson(SUCC, "登录成功")

		}
	}
	self.AjaxJson(FAIL, "请求方式错误")
}

func (self *LoginController) LoginOut() {
	self.DelSession("account")
	self.DelSession("userid")
	self.redirect(beego.URLFor("LoginController.Login"))
}
