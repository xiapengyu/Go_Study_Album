package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	QueryUserList()
}

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_db")
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	log.Println("数据库连接成功...")
	return db
}

func QueryUserList() {
	log.Println("开始查询用户列表")
	db := GetConnection()

	var (
		id      int
		name    string
		address string
		age     int
	)
	rows, err := db.Query("select * from user")
	if err != nil {
		log.Fatal("查询数据失败", err)
	}
	defer rows.Close()

	log.Println("用户列表查询成功")
	for rows.Next() {
		err := rows.Scan(&id, &name, &address, &age)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("id=%d, name=%s, address=%s, age=%d", id, name, address, age)
	}

}
