package controllers

import "github.com/astaxie/beego"

/**
  该控制器处理页面错误请求
*/
type ErrorController struct {
	beego.Controller
}

func (this *ErrorController) Error401() {
	result := Reponse(401, "", "未经授权，请求要求验证身份")
	this.Data["json"] = result
	this.ServeJSON()
}
func (this *ErrorController) Error403() {
	result := Reponse(401, "", "服务器拒绝请求")
	this.Data["json"] = result
	this.ServeJSON()
}
func (this *ErrorController) Error404() {
	result := Reponse(404, "", "很抱歉您访问的地址或者方法不存在")
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *ErrorController) Error500() {
	result := Reponse(500, "", "server error")
	this.Data["json"] = result
	this.ServeJSON()
}
func (this *ErrorController) Error503() {
	result := Reponse(503, "", "服务器目前无法使用（由于超载或停机维护）")
	this.Data["json"] = result
	this.ServeJSON()
}
