package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ipinfo/api/apiMethods"
)

func main() {
	service := apiMethods.Service{}
	router := gin.Default()
	router.POST("/api/getInfo", service.GetInfo)
	err := router.Run(":8081")
	if err != nil {
		fmt.Println("err: ", err)
	}
	//resp, err := http.Get("https://www.artlebedev.ru/country-list/xml/")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//log.Println(string(body))
}
