package config

import (
	"fmt"
	"ginrbac/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/ini.v1"
	"log"
)

var Conf *config //应用配置变量，可以通过它来读取
var Db *gorm.DB  //应用mysql连接池

type config struct {
	AppMode     string
	AppPort     string
	MysqlUser   string
	MysqlPwd    string
	MysqlIpPort string
	MysqlDb     string
	MysqlOption string
}

func init() {
	initConfig()
	initDb()   //调用同级目录下db.go中的方法
	genTable() //生成修改表结构
}
func initConfig() {
	cfg, err := ini.Load("./config.ini")
	if err != nil {
		log.Panic("加载ini错误！")
	}
	Conf = &config{}
	err = cfg.MapTo(Conf)
	if err != nil {
		log.Panic("解析ini文件错误", err)
	}
}

func initDb() {
	var err error
	//gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8&parseTime=True&loc=Local")
	dnsStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", Conf.MysqlUser, Conf.MysqlPwd, Conf.MysqlIpPort, Conf.MysqlDb, Conf.MysqlOption)
	Db, err = gorm.Open("mysql", dnsStr)
	if err != nil {
		log.Panic("连接mysql有错误:", err)
	}
}
func genTable() {
	Db.AutoMigrate(&model.User{})
	Db.AutoMigrate(&model.Role{})
	Db.AutoMigrate(&model.Node{})
}
