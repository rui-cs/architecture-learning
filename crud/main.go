package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rui-cs/architecture-learning/dao"
	"github.com/rui-cs/architecture-learning/service"
	"github.com/rui-cs/architecture-learning/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := initDB()

	server := initWebServer()

	initUser(server, db)

	err := server.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:youre_pass@tcp(localhost:3306)/crud?charset=utf8&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}

	return db
}

func initWebServer() *gin.Engine {
	server := gin.Default()

	return server
}

func initUser(server *gin.Engine, db *gorm.DB) {
	ud := dao.NewUserDao(db)
	us := service.NewUserService(ud)
	c := web.NewUserHandler(us)
	c.RegisterRoutes(server)
}
