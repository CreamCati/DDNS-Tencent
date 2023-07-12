# 动态域名解析服务，用于腾讯DNSPOD

> 需要go 1.20.0 及以上
> 
> 使用go mod download 下载依赖
>
> 需要更改go源，默认大陆地区无法下载依赖，会出现连接超时异常
> 
> `go run DDNS.go` 运行项目

* 前端文件已打包放在templates文件中,如果需要build运行项目，请确保同级目录下有templates文件夹
  > .env文件会自行创建，如果创建失败，请手动生成一个内容为
  > 
  > ID=A
  > 
  > KEY=B
  > 
  > 名称为.env的空白文件

* tips:运行后先设置服务器ip和自己的ID和KEY
* 当前只支持单条记录动态绑定
