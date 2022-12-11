package main

import (
	"github.com/sirupsen/logrus"

	//"service-product/schemas"

	PluginConsul "github.com/go-micro/plugins/v4/registry/consul"
	ConsulAPI "github.com/hashicorp/consul/api"
	"gorm.io/driver/mysql"
	"log"
	"os"
	handlers "service-product/handlers/product"
	"service-product/models"
	"service-product/pkg"
	pb "service-product/proto"
	repositorys "service-product/repositorys/product"
	services "service-product/services/product"

	"go-micro.dev/v4"

	"github.com/go-micro/plugins/v4/server/grpc"
	"gorm.io/gorm"
)

func main() {

	db := SetupDatabase()
	createStudentRepository := repositorys.NewRepositoryCreate(db)
	createStudentService := services.NewServiceCreate(createStudentRepository)

	updateStudentRepository := repositorys.NewRepositoryUpdate(db)
	updateStudentService := services.NewServiceUpdate(updateStudentRepository)

	resultStudentRepository := repositorys.NewRepositoryResult(db)
	resultStudentService := services.NewServiceResult(resultStudentRepository)

	deleteStudentRepository := repositorys.NewRepositoryDelete(db)
	deleteStudentService := services.NewServiceDelete(deleteStudentRepository)

	resultsStudentRepository := repositorys.NewRepositoryResults(db)
	resultsStudentService := services.NewServiceResults(resultsStudentRepository)
	InitProduct := handlers.NewHandlerRPCProduct(resultsStudentService, resultStudentService, createStudentService, updateStudentService, deleteStudentService)

	service := micro.NewService(
		micro.Server(grpc.NewServer()),
		micro.Name("service-product"),
	)
	consulConf := ConsulAPI.DefaultConfig()
	consulConf.Address = "127.0.0.1:8567"
	consulConf.Scheme = "http"

	service.Init(
		micro.Registry(
			PluginConsul.NewRegistry(
				PluginConsul.Config(consulConf))))

	pb.RegisterServiceProductHandlerHandler(service.Server(), InitProduct)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

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
		&models.ModelProduct{},
		//&models.ModelUser{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
