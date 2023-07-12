package utils

import (
	"DDNS/setting"
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	"io/ioutil"
	"log"
	_ "net"
	"net/http"
	"strconv"
	"time"
)

func getPublicIP() (string, error) {
	resp, err := http.Get("https://ipinfo.io/ip")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ipBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ipBytes), nil
}

func DeleteRecord(RecordId string, domain string) string {
	credential := common.NewCredential(
		setting.SecretId,
		setting.SecretKey,
	)
	client, _ := dnspod.NewClient(credential, regions.Guangzhou, profile.NewClientProfile())

	// 填写要删除的记录的 RecordId
	id, _ := strconv.Atoi(RecordId)
	// 删除记录
	deleteRecordRequest := dnspod.NewDeleteRecordRequest()
	deleteRecordRequest.RecordId = common.Uint64Ptr(uint64(id))
	deleteRecordRequest.Domain = common.StringPtr(domain)
	_, err := client.DeleteRecord(deleteRecordRequest)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has occurred: %s", err)
		return err.Error()
	} else if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return "Record with RecordId " + RecordId + " has been deleted"

}
func GetDomainList() string {
	// 配置账号密钥和地域
	credential := common.NewCredential(setting.SecretId,
		setting.SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = setting.TencentApi
	cpf.HttpProfile.ReqTimeout = 5

	// 创建 DNSPod 客户端
	client, _ := dnspod.NewClient(credential, regions.Guangzhou, cpf)

	// 创建请求参数
	request := dnspod.NewDescribeDomainListRequest()

	var jsonArray []map[string]interface{}

	// 发送请求获取域名列表
	response, err := client.DescribeDomainList(request)
	if err != nil {
		return err.Error()
	}
	for i := 0; i < len(response.Response.DomainList); i++ {
		domain := *response.Response.DomainList[i]
		obj := make(map[string]interface{})
		obj["DomainId"] = domain.DomainId
		obj["DomainName"] = domain.Name
		jsonArray = append(jsonArray, obj)
	}
	jsonData, err := json.Marshal(jsonArray)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		return err.Error()
	}
	log.Print(string(jsonData))
	return string(jsonData)
}
func GetDomainRecordList(Domain string) string {
	// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
	// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
	// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
	credential := common.NewCredential(
		setting.SecretId,
		setting.SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = setting.TencentApi
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := dnspod.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := dnspod.NewDescribeRecordListRequest()

	request.Domain = common.StringPtr(Domain)

	// 返回的resp是一个DescribeRecordListResponse的实例，与请求对象对应
	response, err := client.DescribeRecordList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return err.Error()
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	return response.ToJsonString()
}

func CreateRecord(Domain string, SubDomain string, RecordType string, Value string) string {
	// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
	// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
	// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
	credential := common.NewCredential(
		setting.SecretId,
		setting.SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = setting.TencentApi
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := dnspod.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := dnspod.NewCreateRecordRequest()

	// 填入请求参数
	request.Domain = common.StringPtr(Domain)         // 填入你的域名
	request.SubDomain = common.StringPtr(SubDomain)   // 填入你的域名
	request.RecordType = common.StringPtr(RecordType) // 填入记录类型，例如 A、CNAME、MX 等
	request.RecordLine = common.StringPtr("默认")       // 填入线路名称
	request.Value = common.StringPtr(Value)           // 填入记录值
	// 返回的resp是一个CreateDomainResponse的实例，与请求对象对应
	response, err := client.CreateRecord(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return err.Error()
	}
	if err != nil {
		panic(err)
		return err.Error()
	}
	// 输出json格式的字符串回包
	return response.ToJsonString()
}

func ModifyRecord(Domain string, SubDomain string, RecordType string, Value string, RecordId string, RecordLine string) string {
	log.Println(Value)
	// 实例化一个认证对象，入参需要传入腾讯云账户 SecretId 和 SecretKey，此处还需注意密钥对的保密
	// 代码泄露可能会导致 SecretId 和 SecretKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议采用更安全的方式来使用密钥，请参见：https://cloud.tencent.com/document/product/1278/85305
	// 密钥可前往官网控制台 https://console.cloud.tencent.com/cam/capi 进行获取
	credential := common.NewCredential(
		setting.SecretId,
		setting.SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = setting.TencentApi
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := dnspod.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := dnspod.NewModifyRecordRequest()

	// 填入请求参数
	request.Domain = common.StringPtr(Domain)         // 填入你的域名
	request.SubDomain = common.StringPtr(SubDomain)   // 填入你的域名
	request.RecordType = common.StringPtr(RecordType) // 填入记录类型，例如 A、CNAME、MX 等
	request.Value = common.StringPtr(Value)           // 填入记录值
	request.RecordLine = common.StringPtr("默认")       // 填入线路名称

	id, _ := strconv.Atoi(RecordId)
	request.RecordId = common.Uint64Ptr(uint64(id))
	// 返回的resp是一个CreateDomainResponse的实例，与请求对象对应
	response, err := client.ModifyRecord(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		return err.Error()
	}
	if err != nil {
		panic(err)
		return err.Error()
	}
	// 输出json格式的字符串回包
	return response.ToJsonString()
}

var ticker *time.Ticker // 全局定时器变量
var isRunning bool

type Task struct {
	Time     string
	RecordId string
	Domain   string
	Line     string
	Name     string
	Type     string
}

func StartTask(body Task) string {
	if isRunning {
		// 停止当前定时器
		ticker.Stop()
	}
	if body.Time == "0" {
		return "定时任务关闭"
	}
	mtime, _ := strconv.Atoi(body.Time)
	duration := time.Duration(mtime) * time.Minute
	ticker = time.NewTicker(duration)
	// 启动一个 goroutine 来处理定时任务
	go func() {
		isRunning = true
		for {
			// 等待定时器触发
			<-ticker.C
			// 在这里编写定时任务的逻辑
			fmt.Println("定时任务执行" + body.Time)
			ip, _ := getPublicIP()
			log.Print(ModifyRecord(body.Domain, body.Name, body.Type, ip, body.RecordId, body.Line))
		}
	}()

	return "定时任务已开启，每" + body.Time + "分钟执执行一次,再次绑定可重置定时任务"
}
