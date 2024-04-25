package main

import (
	"encoding/json"
	"net/http"
	"sms/io"
	"sms/sdk"
	"sms/utils"

	"github.com/gin-gonic/gin"
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
	sdk.SendSms(body.Phone, code)
	c.JSONP(http.StatusOK, true)
}

func test(c *gin.Context) {
	c.JSONP(http.StatusOK, "test")
}

func main() {

	conf, _ := io.ReadFile()
	port := utils.ParsePort(conf["port"].(string))

	r := gin.Default()
	r.GET("/test", test)
	r.POST("/sms", sendCode)

	err := r.Run(port)
	if err != nil {
		return
	}

}
