package main

import (
	"Ecommerce/components/appctx"
	"Ecommerce/config"
	_ "Ecommerce/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.

// @host      localhost:3007
// @BasePath  /api/v1
func main() {
	log.Println("server start")
	dsn := os.Getenv("MY_SQL_CONN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Can not connect to db")
	}
	db = db.Debug()
	db = config.AutoMigrateDB(db)
	appctx := appctx.NewAppContext(db)

	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	RouteConfig(appctx, r)
	r.Run(":3007")
}
