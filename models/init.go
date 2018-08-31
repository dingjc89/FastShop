/**
 *
 *@author steve dingjc89@126.com
 *2018/8/15
 *@version V1.0
 *@copyright steve
 */
package models

import (
	"net/url"

	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func Init() {
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpass := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	dsn := dbuser + ":" + dbpass + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	fmt.Println(dsn)
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(
		new(User),
		new(Role),
		new(Permissions),
		new(Logs),
	)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	orm.RunCommand()

}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}

// 过滤器处理
func Filter(filter map[string][]interface{}) map[string][]interface{} {
	f := make(map[string][]interface{}, len(filter))
	for key, slice := range filter {
		switch len(slice) {
		case 1:
			f[key] = slice[0:]
		default:
			exp := slice[0]
			key += "__" + exp.(string)
			f[key] = slice[1:]
		}
	}
	return f
}
