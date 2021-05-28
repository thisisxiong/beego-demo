package init

import (
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

	username, _ := config.String("mysql::username")
	password, _ := config.String("mysql::password")
	host, _ := config.String("mysql::host")
	port, _ := config.String("mysql::port")
	dbname, _ := config.String("mysql::dbname")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", username, password, host, port, dbname)
	dataSource := orm.RegisterDataBase("default", "mysql", dsn)
	if dataSource != nil {
		logs.Info(fmt.Sprintf("msyql connect error :%s", dataSource))
	}
}
