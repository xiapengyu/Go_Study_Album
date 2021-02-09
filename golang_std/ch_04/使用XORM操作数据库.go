package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

func main() {

}

func getMySQLConnection() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123456@(127.0.0.1:3306)/go_db?charset=utf8")
	if err != nil {
		log.Panicln("数据库连接失败", err)
	}
	//连接池设置
	engine.SetMaxOpenConns(20)
	engine.SetMaxIdleConns(10)
	engine.SetConnMaxLifetime(30 * time.Second)
	//打印SQL
	engine.ShowSQL(true)
	log.Println("数据库连接成功")

	//设置表名与字段名的映射规则
	engine.SetColumnMapper(names.SnakeMapper{})	//结构体驼峰，数据库字段下划线分割
	engine.SetTableMapper(names.SnakeMapper{})	//结构体驼峰，数据库表名下划线分割
	//engine.SetTableMapper(names.NewPrefixMapper(names.SnakeMapper{}, "tb_"))	//前缀映射
	//engine.SetTableMapper(names.NewSuffixMapper(names.SnakeMapper{}, "_tb"))	//后缀映射
	return engine
}

/**
创建引擎组
*/
func createEngineGroup() *xorm.EngineGroup {
	urls := []string{
		"root:123456@(127.0.0.1:3306)/go_db?charset=utf8", // 第一个默认是master
		"root:123456@(127.0.0.1:3306)/go_db?charset=utf8", // 第二个开始都是slave
		"root:123456@(127.0.0.1:3306)/go_db?charset=utf8",
	}
	//所有引擎使用相同的driverName，只能连接一种数据库
	engineGroup, err := xorm.NewEngineGroup("mysql", urls, xorm.RandomPolicy()) //负载均衡策略：随机访问，权重随机访问，轮训访问，权重轮训访问，最小连接数访问
	if err != nil {
		log.Panicln("创建引擎组失败", err)
	}

	engineGroup.SetMaxOpenConns(20)
	engineGroup.SetMaxIdleConns(10)
	engineGroup.SetConnMaxLifetime(30 * time.Second)
	engineGroup.ShowSQL(true)
	return engineGroup
}

/**
创建引擎组
*/
func createEngineGroup1() *xorm.EngineGroup {
	//每个引擎都可以设置不同的driverName，通过设置不同的driverName，可以连接不同的数据库
	master0, _ := xorm.NewEngine("mysql", "root:123456@(127.0.0.1:3306)/go_db?charset=utf8")
	master1, _ := xorm.NewEngine("mysql", "root:123456@(127.0.0.1:3306)/go_db?charset=utf8")
	slave0, _ := xorm.NewEngine("mysql", "root:123456@(127.0.0.1:3306)/go_db?charset=utf8")
	slave1, _ := xorm.NewEngine("mysql", "root:123456@(127.0.0.1:3306)/go_db?charset=utf8")
	slave2, _ := xorm.NewEngine("mysql", "root:123456@(127.0.0.1:3306)/go_db?charset=utf8")

	masters := []*xorm.Engine{master0, master1}
	slaves := []*xorm.Engine{slave0, slave1, slave2}
	engineGroup, err := xorm.NewEngineGroup(masters, slaves, xorm.RandomPolicy())
	if err != nil {
		log.Panicln("创建引擎组失败", err)
	}

	engineGroup.SetMaxOpenConns(20)
	engineGroup.SetMaxIdleConns(10)
	engineGroup.SetConnMaxLifetime(30 * time.Second)
	engineGroup.ShowSQL(true)
	return engineGroup
}
