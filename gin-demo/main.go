package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
)

type UserSelectConditionReq struct {
	Account string `json:"account"`
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		client := resty.New()
		response, err := client.R().
			SetBody(&UserSelectConditionReq{Account: "esladmin"}).
			Post("http://192.168.100.148:9901/zk-user-center-core/user/getUserByCondition")
		if err != nil {
			log.Fatalln(err)
			return
		}
		log.Println(response)
		c.JSON(http.StatusOK, string(response.Body()))
	})

	r.Run()
}
