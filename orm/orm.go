package orm

import "C"
import (
	"fmt"
	clog "log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/strconv/valyria/config"
	"github.com/strconv/valyria/log"
)

var db *gorm.DB

func Init(c *config.Conf) {
	if c.Mysql.DBName == "" {
		return
	}
	var err error
	db, err = gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			c.Mysql.UserName, c.Mysql.Password, c.Mysql.Addr, c.Mysql.Port, c.Mysql.DBName))
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	// 关闭复数表名，如果设置为true，`User`表的表名就会是`user`，而不是`users`
	db.SingularTable(true)
	// 启用Logger，显示详细日志
	db.LogMode(true)
	db.SetLogger(log.NewGormLogger())
	//连接池
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(200)
	clog.Print("mysql init done! ")
}

//get db
func GetDB() *gorm.DB {
	return db
}
