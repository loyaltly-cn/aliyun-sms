package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_server/io"
	"go_server/sms"
	"go_server/utils"
	"net/http"
	"reflect"
)

type Rep struct {
	Code  string `json:"code"`
	Phone string `json:"phone"`
}

func sendCode(c *gin.Context) {
	b, _ := c.GetRawData()
	body := Rep{}
	_ = json.Unmarshal(b, &body)
	code := utils.ParseCode(body.Code)
	sms.SendSms(body.Phone, code)
	c.JSONP(http.StatusOK, true)
}

func test(c *gin.Context) {
	c.JSONP(http.StatusOK, "test")
}

func main() {

	conf, _ := io.ReadFile()
	port := utils.ParseProt(conf["port"].(string))

	fmt.Println(reflect.TypeOf(conf["port"]))
	r := gin.Default()
	r.GET("/test", test)
	r.POST("/sendCode", sendCode)

	err := r.Run(port)
	if err != nil {
		return
	}

}
