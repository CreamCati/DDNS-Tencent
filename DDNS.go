package main

import (
	"DDNS/controller"
	"DDNS/web"
	"github.com/gin-gonic/gin"
	_ "net"
	"net/http"
)

func main() {
	//log.Print(utils.ModifyRecord("klistem.eu.org", "@", "A", "2.3.4.5", "1427840446"))
	route := gin.Default()
	route.Use(Cors())
	route.Static("/asset", "templates/asset")
	route.LoadHTMLGlob("templates/**/*")
	// 2.绑定路由规则，执行的函数

	route.GET("/", web.Home)
	route.GET("/domain/list", controller.DomainList)
	route.POST("/domain/info", controller.DomainInfo)
	route.POST("/domain/task", controller.DomainTask)
	route.POST("/domain/modify", controller.DomainModify)
	route.POST("/domain/delete", controller.DomainDelete)
	route.POST("/domain/create", controller.DomainCreate)
	route.POST("/setting/auth", controller.SettingAuth)
	route.Run(":" + "8325")

	//ip, err := getPublicIP()
	//if err != nil {
	//	fmt.Println("Failed to get public IP:", err)
	//	return
	//}
	//
	//fmt.Println("Public IP:", ip)
	//RecordList := getList()
	//log.Print(RecordList)
	//deleteRecord(1525724344)
	//log.Print(modifyRecord("klistem.eu.org", "@", "A", "1.2.3.24"))

}

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, DELETE, PUT, HEAD")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
