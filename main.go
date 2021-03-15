package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log-exploration-api/pkg/configuration"
	"log-exploration-api/pkg/elastic"
	"log-exploration-api/pkg/logscontroller"
)

func main() {
	router := gin.Default()
	appConf := configuration.ParseArgs()
	repository,err := elastic.NewElasticRepository(appConf)
	fmt.Println(repository)
	if(err!=nil){
		fmt.Println(err)
		return
	}
	logscontroller.NewLogsController(repository, router)
	router.Run()
}
