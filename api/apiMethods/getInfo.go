package apiMethods

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ipinfo/methods"
)

func (service *Service) GetInfo(c *gin.Context) {
	var a methods.Data
	method := methods.Service{}

	err := c.BindJSON(&a)

	if err != nil {
		fmt.Println("err: ", err)
	}

	for i := 0; i < len(a.Data); i++ {
		method.GetInfo(a.Data[i].Ip)
	}

}
