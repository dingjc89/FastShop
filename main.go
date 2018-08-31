package main

import (
	"quickstart/models"
	_ "quickstart/routers"

	"quickstart/libs"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init() {
	// 初始化数据模型
	models.Init()
	// 模版函数映射
	beego.AddFuncMap("Permission", libs.Permission)
}

func main() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"project.log"}`)
	logs.EnableFuncCallDepth(true) // 设置输出日志文件，行号
	logs.Async()                   // 异步输出日志
	beego.Run()
}
