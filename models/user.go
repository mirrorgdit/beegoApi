package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64
	Username string    `orm:"size(20)"`
	Avatar   string    `orm:"size(255)"`
	Phone    string    `orm:"size(20)"`
	Password string    `orm:"size(32)"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
}

type UserForm struct {
	UserName string `valid:"Required;MinSize(4);MaxSize(20)"` // Name     不能为空并且最小长度是4 最大长度是20
	PassWord string `valid:"Required;MinSize(6);MaxSize(20)"` // PassWord 不能为空并且最小长度是6 最大长度是20
	Phone    string `valid:"Mobile"`                          // Mobile 必须为正确的手机号
}

//获取模型表名
func (m *User) TableName() string {
	return "user"
}

//获取全部用户
func (m *User) GetAllUser() []*User {
	info := User{}
	list := make([]*User, 0)
	info.Query().All(&list)
	return list
}

//添加用户
func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

//删除用户
func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

//修改用户信息
func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Read(fields ...string) (int64, error) {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return 0, err
	}
	return m.Id, nil
}

//通过ID查询单个用户信息
func (m *User) GetUserById(uid int64) (*User, error) {
	info := &User{}
	err := info.Query().Filter("Id", uid).One(info)
	return info, err
}

//通过用户名或者单个用户信息
func (m *User) GetUserByName(name string) (*User, error) {
	info := &User{}
	err := info.Query().Filter("Username", name).One(info)
	return info, err
}

//查询方法
func (m *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
