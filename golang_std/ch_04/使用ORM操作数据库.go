package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Like struct {
	ID        int    `gorm:"primary_key"`
	Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua        string `gorm:"type:varchar(256);not null;"`
	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
	Hash      uint64 `gorm:"unique_index:hash_idx;"`
	CreatedAt time.Time
}

func main() {
	//createTable()
	//save()
	search()
}

func getConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/go_db?charset=utf8")
	if err != nil {
		log.Panicln("数据库连接失败", err)
	} else {
		log.Println("数据库连接成功")
	}
	db.DB().SetMaxOpenConns(20)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(30)
	db.SingularTable(true) //使用结构体建表时不会生成后缀S
	db.LogMode(true)       //打印sql语句
	return db
}

func createTable() {
	log.Println("根据结构体创建数据表")
	db := getConnect()
	if !db.HasTable(&Like{}) {
		err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Like{}).Error
		if err != nil {
			log.Panicln("数据表创建失败", err)
		} else {
			log.Println("数据表创建成功")
		}
	} else {
		log.Println("数据表已经存在")
	}
}

func save() {
	db := getConnect()
	defer db.Close()
	like := Like{2, "127.0.0.1", "ua", "点赞关注", 12345678901234567891, time.Now()}
	log.Println("插入数据", like)
	db.Create(like)
}

func search() {
	db := getConnect()
	defer db.Close()

	//简单sql查询
	rows, err := db.Select([]string{"id", "ip"}).Find(&Like{}).Rows()
	if err != nil {
		log.Panicln("查询失败", err)
	}
	for rows.Next() {
		s := Like{}
		if err := rows.Scan(&s); err != nil {
			return
		}
		log.Println(s)
	}
}
