package main

import (
	"fmt"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"os"
	"service-product/enitites"
	"service-product/pkg"
	"service-product/routes"
	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	app := SetupRouter()
	logrus.Fatal(app.Run(":" + pkg.GodotEnv("GO_PORT")))

}

func SetupRouter() *gin.Engine {
	db := SetupDatabase()
	app := gin.Default()

	fmt.Println("GO_ENV main ", pkg.GodotEnv("GO_ENV"))
	if pkg.GodotEnv("GO_ENV") != "production" && pkg.GodotEnv("GO_ENV") != "test" {
		gin.SetMode(gin.DebugMode)
	} else if pkg.GodotEnv("GO_ENV") == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)

	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	app.Use(helmet.Default())
	app.Use(gzip.Gzip(gzip.BestCompression))
	routes.InitProductRoutes(db, app)

	return app
}

func SetupDatabase() *gorm.DB {
	urldb := pkg.GodotEnv("DATABASE_URI")
	//fmt.Println("urldb ",urldb)
	db, err := gorm.Open(mysql.Open(urldb), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connect into Database Failed")
		logrus.Fatal(err.Error())
	}

	if os.Getenv("GO_ENV") != "production" {
		logrus.Info("Connect into Database Successfully")
	}

	err = db.AutoMigrate(
		&enitites.EntityProduct{},
		//&models.ModelUser{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
