package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 连接MySQL数据库
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/testdata")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 创建表格
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT NOT NULL AUTO_INCREMENT, name VARCHAR(50) NOT NULL, age INT NOT NULL, PRIMARY KEY (id))")
	if err != nil {
		panic(err.Error())
	}

	// 插入数据
	stmt, err := db.Prepare("INSERT INTO users(name, age) VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec("Tom", 25)
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec("Jerry", 28)
	if err != nil {
		panic(err.Error())
	}

	// 查询数据
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
}
