/**
 *
 *@author steve dingjc89@126.com
 *2018/8/16
 *@version V1.0
 *@copyright steve
 */
package controllers

type HomeController struct {
	BaseController
}

func (self *HomeController) Index() {
	self.display("home/index")
}
