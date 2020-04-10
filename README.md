# BeegoApi Demo

### beego restful api, jwt过滤器

## 使用

* https://github.com/mirrorgdit/beegoApi.git
* bee run

## 测试路由
### 获取所有用户数据
* GET 获取所有用户数据 http://localhost:8080/api/v1/user

### 用户注册
* POST 用户注册 http://localhost:8080/api/v1/user/1/register

### 认证测试
* POST 用户登陆 http://localhost:8080/api/v1/user/login

### 认证测试
* jwt用户验证测试，需要传入header - Authorization参数  http://localhost:8080/api/v1/user/auth

### 其他具体见路由文件
* ....
