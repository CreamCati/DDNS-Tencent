package controller

import (
	"DDNS/setting"
	"DDNS/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
)

func SettingAuth(context *gin.Context) {
	type Auth struct {
		SecretId  string
		SecretKey string
	}
	var body Auth
	err := context.Bind(&body)
	if err != nil {
		log.Println(err)
		return
	}

	filePath := ".env"

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("无法读取文件:", err)
		return
	}
	newContent := string(content)

	newContent = strings.ReplaceAll(newContent, "ID="+setting.SecretId, "ID="+body.SecretId)

	newContent = strings.ReplaceAll(newContent, "KEY="+setting.SecretKey, "KEY="+body.SecretKey)

	err = os.WriteFile(filePath, []byte(newContent), 0644)
	if err != nil {
		fmt.Println("无法写入文件:", err)
		context.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "无法写入文件",
		})
		return
	}
	log.Print("修改成功")
	setting.SecretId = body.SecretId
	setting.SecretKey = body.SecretKey
	log.Print(body.SecretKey)
	context.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "settingOk,new ID and Key:" + body.SecretId + "----------" + setting.SecretKey,
	})

}
func DomainCreate(context *gin.Context) {
	type Domain struct {
		Domain string
		Name   string
		Type   string
		Value  string
	}
	var body Domain
	err := context.Bind(&body)
	if err != nil {
		log.Println(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": utils.CreateRecord(body.Domain, body.Name, body.Type, body.Value),
	})
}

func DomainDelete(context *gin.Context) {
	type Domain struct {
		RecordId string
		Domain   string
	}
	var body Domain
	err := context.Bind(&body)
	if err != nil {
		log.Println(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": utils.DeleteRecord(body.RecordId, body.Domain),
	})
}
func DomainModify(context *gin.Context) {
	type Domain struct {
		Domain     string
		SubDomain  string
		RecordType string
		Value      string
		RecordId   string
	}
	var body Domain
	err := context.Bind(&body)
	if err != nil {
		log.Println(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
	})
}

func DomainTask(context *gin.Context) {

	var body utils.Task
	err := context.Bind(&body)

	if err != nil {
		log.Println(err)
		return
	}

	msg := utils.StartTask(body)
	context.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": msg,
	})
}
func Create(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": gin.H{
			"cpu": "",
		},
	})
}

func Delete(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": gin.H{
			"cpu": "",
		},
	})
}

func DomainList(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": utils.GetDomainList(),
	})
}

func DomainInfo(context *gin.Context) {

	type Domain struct {
		Name string
	}
	var body Domain
	err := context.Bind(&body)
	if err != nil {
		log.Println(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": utils.GetDomainRecordList(body.Name),
	})
}
