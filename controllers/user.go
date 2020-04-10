package controllers

import (
	"beegoApi/models"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	BaseController
}

// @Title 用户登陆
// @Description 用户登陆 http://localhost:8080/api/v1/user/1/update
// @Param   username
// @Param   password
// @Success 2000
// @Failure 4001 User not found
// @router / [post]
func (this *UserController) Login() {
	result := DataResponse{}
	userForm := models.UserForm{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &userForm)
	valid := validation.Validation{}
	b, err := valid.Valid(&userForm)
	if err != nil {
		return
	}
	if !b {
		for _, err := range valid.Errors {
			result = Reponse(4001, userForm, err.Key+err.Message)
			this.Data["json"] = result
			this.ServeJSON()
			return
		}
	}
	userMod := &models.User{Username: userForm.UserName, Password: userForm.PassWord}
	uId, err := userMod.Read("Username", "Password")
	if err == nil {
		tokenStr, err := this.GenToken(uId)
		if err != nil {
			beego.Debug(err.Error())
		}
		result = Reponse(2000, tokenStr, "")
	} else {
		result = Reponse(4001, "", "username or password error")
	}
	this.Data["json"] = result
	this.ServeJSON()
}

// @Title token验证
// @Description jwt用户验证测试，需要传入header - auth参数  http://localhost:8080/api/v1/user/auth
// @Param   header key: auth
// @Success 2000
// @Failure 4004 User not found
// @router / [get]
func (this *UserController) Auth() {
	result := DataResponse{}
	_, isValid, err := this.ValidToken()
	if isValid {
		result = Reponse(2000, "", "auth success")
	} else {
		result = Reponse(4001, "", err.Error())
	}
	this.Data["json"] = result
	this.ServeJSON()
}

// @Title 获取所有用户数据
// @Description 获取所有用户数据 http://localhost:8080/api/v1/user
// @Success 2000
// @Failure 4004 User not found
// @router / [get]
func (this *UserController) GetAll() {
	result := DataResponse{}
	userModel := models.User{}
	list := userModel.GetAllUser()
	if len(list) == 0 {
		result = Reponse(4004, "", "Data No Found")
	} else {
		result = Reponse(2000, list, "OK")
	}
	this.Data["json"] = result
	this.ServeJSON()

}

// @Title 用户注册
// @Description 用户注册 http://localhost:8080/api/v1/user/1/register
// @Param   username
// @Param   password
// @Success 2000
// @Failure 4001 User not register
// @router / [post]
func (this *UserController) Register() {
	valid := validation.Validation{}
	result := DataResponse{}
	userForm := models.UserForm{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &userForm)

	b, err := valid.Valid(&userForm)
	if err != nil {
		return
	}
	if !b {
		// validation does not pass
		for _, err := range valid.Errors {
			result = Reponse(4000, userForm, err.Key+err.Message)
			this.Data["json"] = result
			this.ServeJSON()
			return
		}
	}

	userModel := models.User{}
	if u, err := userModel.GetUserByName(userForm.UserName); err == nil && u.Id != 0 {
		result = Reponse(4000, userForm, "帐号已经存在!")
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	userModel.Username = userForm.UserName
	userModel.Password = userForm.PassWord
	userModel.Phone = userForm.Phone
	if err := userModel.Insert(); err != nil {
		result = Reponse(4000, "", "username or password error")
	} else {
		result = Reponse(2000, "", "注册成功")
	}
	this.Data["json"] = result
	this.ServeJSON()
}

// @Title 获取某个用户的信息
// @Description 获取所有用户数据 http://localhost:8080/api/v1/user/1
// @Success 2000
// @Failure 4004 User not found
// @router / [get]
func (this *UserController) GetOne() {
	result := DataResponse{}
	uid, _ := this.GetInt64(":id")
	userModel := models.User{Id: uid}
	if userInfo, err := userModel.GetUserById(uid); err != nil {
		result = Reponse(4004, userInfo, "get user fail")
	} else {
		result = Reponse(2000, userInfo, "get user success")
	}
	this.Data["json"] = result
	this.ServeJSON()
}

// @Title 删除某个用户信息
// @Description 删除某个用户信息 http://localhost:8080/api/v1/user/1/del
// @Success 2000
// @Failure 4004 del user err
// @router / [post]
func (this *UserController) Delete() {
	result := DataResponse{}
	uid, _ := this.GetInt64(":id")
	userMod := models.User{Id: uid}
	if err := userMod.Delete(); err != nil {
		result = Reponse(4006, "", "del user fail")
	} else {
		result = Reponse(2000, "", "del user success")
	}
	this.Data["json"] = result
	this.ServeJSON()
}

// @Title 更新某用户信息
// @Description 获取所有用户数据 http://localhost:8080/api/v1/user/1/update
// @Success 2000
// @Failure 4004 User not found
// @router / [post]
func (this *UserController) Update() {

	valid := validation.Validation{}
	result := DataResponse{}
	userForm := models.UserForm{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &userForm)
	b, err := valid.Valid(&userForm)
	if err != nil {
		return
	}
	if !b {
		// validation does not pass
		for _, err := range valid.Errors {
			result = Reponse(4000, userForm, err.Key+err.Message)
			this.Data["json"] = result
			this.ServeJSON()
			return
		}
	}

	uid, _ := this.GetInt64(":id")
	userModel := models.User{Id: uid}
	userModel.Phone = userForm.Phone
	if err := userModel.Update("Phone"); err != nil {
		result = Reponse(4005, "", "update fail")
	} else {
		result = Reponse(2000, "", "update success")
	}
	this.Data["json"] = result
	this.ServeJSON()
}
