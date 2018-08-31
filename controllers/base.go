/**
 *
 *@author steve dingjc89@126.com
 *2018/8/15
 *@version V1.0
 *@copyright steve
 */
package controllers

import (
	"strings"

	"quickstart/models"

	"github.com/astaxie/beego"
)

const (
	FAIL = -1
	SUCC = 0
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	userId         int
	userName       string
	loginName      string
	pageSize       int
}

func (self *BaseController) Prepare() {
	self.pageSize = 10
	controllerName, actionName := self.GetControllerAndAction()
	self.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	self.actionName = strings.ToLower(actionName)
	self.Data["siteName"] = "用户管理后台"
	self.Auth()
}

// 用户验证
func (self *BaseController) Auth() {
	var userId int
	if self.GetSession("userid") != nil {
		userId = self.GetSession("userid").(int)
	}
	userM := new(models.User)
	if userId > 0 {
		user, err := userM.GetUserById(userId)
		if err == nil {
			self.userName = user.Name
			self.loginName = user.Account
			self.userId = user.Id

		}
		self.Data["userid"] = userId
		self.Menu(user.Id)
	}
	if self.userId == 0 && (self.controllerName != "login" && self.actionName != "loginin") {
		self.redirect(beego.URLFor("LoginController.Login"))
	}
}

// 左菜单栏
func (self *BaseController) Menu(userId int) {
	filter := models.Filter(map[string][]interface{}{"is_show": {1}})
	permission := new(models.Permissions)
	menu := permission.GetPermissionList(filter)
	list1 := make([]map[string]interface{}, len(menu))
	list2 := make([]map[string]interface{}, len(menu))
	i, j := 0, 0
	for _, row := range menu {
		if row.Pid == 0 {
			r := map[string]interface{}{
				"Id":   int(row.Id),
				"Name": row.Name,
				"Pid":  int(row.Pid),
				"Icon": row.Img,
				"Url":  row.Url,
				"Pn":   row.Pn,
			}
			list1[i] = r
			i++
		} else {
			r := map[string]interface{}{
				"Id":   int(row.Id),
				"Name": row.Name,
				"Pid":  int(row.Pid),
				"Icon": row.Img,
				"Url":  row.Url,
				"Pn":   row.Pn,
			}
			list2[j] = r
			j++
		}
	}
	self.Data["Menu1"] = list1[:i]
	self.Data["Menu2"] = list2[:j]
}

// ajax请求 json数据返回
func (self *BaseController) AjaxJson(status int, msg string) {
	result := make(map[string]interface{})
	result["status"] = status
	result["msg"] = msg
	self.Data["json"] = result
	self.ServeJSON()
	self.StopRun()
}

func (self *BaseController) AjaxList(status int, msg string, count int64, data interface{}) {
	result := make(map[string]interface{})
	result["code"] = status
	result["msg"] = msg
	result["count"] = count
	result["data"] = data
	self.Data["json"] = result
	self.ServeJSON()
	self.StopRun()
}

// 判断是否post请求
func (self *BaseController) isPost() bool {
	return self.Ctx.Request.Method == "POST"
}

// 获取客户端IP
func (self *BaseController) getClientIP() string {
	s := strings.Split(self.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

// 模版渲染
func (self *BaseController) display(tpl ...interface{}) {
	var tplName string
	// 模版路径计算
	// 有参数，第一个参数为模板
	// 没有参数，默认以控制器为目录，方法为文件名
	if len(tpl) > 0 {
		tplName = tpl[0].(string) + ".html"
	} else {
		tplName = self.controllerName + "/" + self.actionName + ".html"
	}
	if len(tpl) > 1 && tpl[1].(bool) == true {
		self.Layout = "public/layout.html"
	}
	self.TplName = tplName
}

// 重定向
func (self *BaseController) redirect(url string) {
	self.Redirect(url, 302)
	self.StopRun()
}
