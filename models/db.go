package models

import (
	"database/sql"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
)

var dbConf config.Configer

func init() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		logs.Info(fmt.Sprintf("mysql register driver error :%s", err))
	}
	//加载数据库配置
	dbConf, err = config.NewConfig("ini", "conf/mysql.conf")
	if err != nil {
		logs.Error("mysql init conf err")
	}
}

func Open(model BaseDb) orm.Ormer {
	dbname := model.DbName()
	var dbcon *sql.DB
	dbcon, err := orm.GetDB(dbname)
	if err != nil {
		mysqlconf := getConf(dbname)
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", mysqlconf.username, mysqlconf.password, mysqlconf.host, mysqlconf.port, dbname)

		logs.Notice(fmt.Sprintf("###获取不到数据库连接 创建数据库:%s 连接###", dbname))
		dataSource := orm.RegisterDataBase(dbname, "mysql", dsn)
		if dataSource != nil {
			logs.Info(fmt.Sprintf("msyql connect error :%s", dataSource))
		}

		return orm.NewOrmUsingDB(dbname)
	}
	//保活
	dbcon.Ping()
	return orm.NewOrmUsingDB(dbname)
}

//数据库配置结构
type Mysqlconf struct {
	host     string
	port     string
	username string
	password string
}

//获取数据库配置
func getConf(name string) Mysqlconf {
	var conf Mysqlconf
	conf.username, _ = dbConf.String(name + "::username")
	conf.password, _ = dbConf.String(name + "::password")
	conf.host, _ = dbConf.String(name + "::host")
	conf.port, _ = dbConf.String(name + "::port")

	return conf
}
