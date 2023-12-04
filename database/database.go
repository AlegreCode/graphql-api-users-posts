package database

import (
	"github.com/alegrecode/graphql-api-users-posts/graph/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conexion() {
	dsn := "root:@tcp(127.0.0.1:3306)/graphql_users_posts?charset=utf8mb4&parseTime=True&loc=Local"
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	DB.AutoMigrate(&model.User{}, &model.Post{})

}
