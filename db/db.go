package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/xormplus/core"
	"github.com/xormplus/xorm"

	"log"
)

var MasterDB *xorm.Engine
var Mysqldb *sql.DB

// Setup initializes the database instance
func Setup() {
	var err error

	MasterDB, err = xorm.NewEngine("mysql", "root:123456@tcp(localhost)/testdata?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	MasterDB.SetMaxIdleConns(3)
	MasterDB.SetMaxOpenConns(15)

	Mysqldb, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/testdata?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

}
