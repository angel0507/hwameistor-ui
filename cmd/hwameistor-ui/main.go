package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"os"

	routers "github.com/hwameistor/hwameistor-ui/server/router"
)

func main() {

	//InitConfig()

	r := gin.Default()
	r = routers.CollectRoute(r)

	//port := viper.GetString("server.port")
	port := "8081"
	fmt.Println("main port %v", port)

	gin.SetMode(gin.ReleaseMode)
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}

func InitConfig() {
	fmt.Printf("InitConfig start ... ")
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	fmt.Println("InitConfig workDir = %v", workDir)
	viper.AddConfigPath(workDir + "/server/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}
}
