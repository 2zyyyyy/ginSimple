package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// gorm  mysql

// User 定义模型
type User struct {
	gorm.Model
	Name     *string `gorm:"default:'空'"`
	Age      *int    `gorm:"default:999"`
	Birthday time.Time
}

var name = "万里"
var age = 18

func mysqlInit() *gorm.DB {
	dsn := "root:数据库连接密码@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("mysql 数据库连接成功！")
	return db
}

// mysql 插入数据
func insert() {
	// 自动迁移
	db := mysqlInit()
	err := db.AutoMigrate(&User{})
	if err != nil {
		fmt.Printf("自动迁移失败，err:%s\n", err)
	}
	db = mysqlInit()
	// 写入数据：实例化 user 对象写入数据
	user := User{
		Name:     &name,
		Age:      &age,
		Birthday: time.Now(),
	}
	// 通过数据的指针来创建
	db.Debug().Create(&user)
	//fmt.Printf("result", result)
}

// mysql 查询数据
func query() {
	var users []User
	// 指定结构体查询字段
	db := mysqlInit()
	db.Debug().Where(&User{Name: &name}, "name").Find(&users) //  SELECT * FROM `users` WHERE
	// `users`.`name` = '万里' AND `users`.`deleted_at` IS NULL
}

// mysql 更行数据
func update() {
	db := mysqlInit()
	newAge := 20
	// 条件更新单列
	db.Debug().Model(&User{}).Where("name = ?", &name).Update("age", &newAge) //  UPDATE `users` SET `
	// age`=20,`updated_at`='2022-05-02 13:20:54.373' WHERE name = '万里' AND `users`.`deleted_at` IS NULL
}

// mysql 删除数据
func deleted() {
	db := mysqlInit()
	// 软删除
	db.Debug().Where("age = ?", &age).Delete(&User{}) //  UPDATE `users` SET `deleted_at`='2022-05-02
	// 13:33:03.749' WHERE age = 18 AND `users`.`deleted_at` IS NULL

}

func main() {
	insert()
	query()
	update()
	deleted()
}
