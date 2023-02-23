package main

import (
	"fmt"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	ConsulAPI "github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
	handlers "service-product/controller/product"
	"service-product/enitites"
	"service-product/pkg"
	pb "service-product/proto/product"
	repositorys "service-product/repositorys/product"
	"service-product/routes"
	services "service-product/services/product"
	"strconv"
)

func main() {

	RESTPort, _ := strconv.Atoi(pkg.GodotEnv("REST_PORT"))
	GRPCPort, _ := strconv.Atoi(pkg.GodotEnv("GRPC_PORT"))
	app := SetupRouter(RESTPort)

	go SetupGRPC(GRPCPort)
	app.Run(fmt.Sprintf(":%v", RESTPort))

}

func SetupConsul(service_name string, port int) {
	consulConf := ConsulAPI.DefaultConfig()
	consulConf.Address = "127.0.0.1:8500"
	consulConf.Scheme = "http"

	client, err := ConsulAPI.NewClient(ConsulAPI.DefaultConfig())
	if err != nil {
		panic(err)
	}

	registration := &ConsulAPI.AgentServiceRegistration{
		Name: service_name,
		Port: port,
	}

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
}

func SetupGRPC(port int) error {
	SetupConsul(pkg.GodotEnv("SERVICE_NAME")+"-grpc", port)

	db := SetupDatabase()
	productRepository := repositorys.NewProductRepositorySQL(db)
	productService := services.NewServiceProduct(productRepository)

	InitProduct := handlers.NewControllerProductRPC(productService)

	s := grpc.NewServer()
	pb.RegisterServiceProductHandlerServer(s, InitProduct)

	log.Println("Starting RPC server at", port)
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("could not listen to %s: %v", port, err)
	}

	return s.Serve(l)
}

func SetupRouter(port int) *gin.Engine {

	SetupConsul(pkg.GodotEnv("SERVICE_NAME")+"-rest", port)

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
