package models

import (
	"database/sql"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		logs.Info(fmt.Sprintf("mysql register driver error :%s", err))
	}
}

func Open(model BaseDb) orm.Ormer {
	dbname := model.DbName()
	var dbcon *sql.DB
	dbcon, err := orm.GetDB(dbname)
	if err != nil {
		//todo 根据dbname获取不同的数据库配置
		username, _ := config.String("mysql::username")
		password, _ := config.String("mysql::password")
		host, _ := config.String("mysql::host")
		port, _ := config.String("mysql::port")
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", username, password, host, port, dbname)

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
