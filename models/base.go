package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	//连接名称
	dbAlias := beego.AppConfig.String("db_alias")
	//数据库名称
	dbName := beego.AppConfig.String("db_name")
	//数据库连接用户名
	dbUser := beego.AppConfig.String("db_user")
	//数据库连接用户名
	dbPwd := beego.AppConfig.String("db_pwd")
	//数据库IP（域名）
	dbHost := beego.AppConfig.String("db_host")
	//数据库端口
	dbPort := beego.AppConfig.String("db_port")
	//数据库编码
	dbCharset := beego.AppConfig.String("db_charset")
	//数据库表前缀
	tbPrefix := beego.AppConfig.String("tb_prefix")
	//注册模型
	orm.RegisterModelWithPrefix(tbPrefix, new(User))
	orm.RegisterDataBase(dbAlias, "mysql", dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset="+dbCharset+"&loc=Asia%2FShanghai")

	orm.RunSyncdb(dbAlias, false, true)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

}
