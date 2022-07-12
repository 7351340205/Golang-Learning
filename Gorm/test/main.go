package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID   int64
	Age  int64
	Name string
	gorm.Model
}

func main() {
	db, sqlDB, _ := connect()
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(sqlDB)

	//db.Migrator().CreateTable(&User{})
	result := db.Find(&User{})
	user := result.Statement.Dest.(*User)
	fmt.Println(user.Age)
	fmt.Println(user.ID)
	fmt.Println(user.Name)
	fmt.Println(user.Model)

}

func connect() (db *gorm.DB, sqlDB *sql.DB, err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, _ = db.DB()
	if err != nil {
		fmt.Printf(err.Error())
		defer sqlDB.Close()
	} else {
		fmt.Printf("数据库连接成功！\n")
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}
	return
}
