package main

import (
	"fmt"
	"github.com/gorcon/rcon"
)

func main() {
	// 创建一个 RCON 连接
	conn, err := rcon.Dial("kursharp.cn:25575", "xxwgyp")
	if err != nil {
		fmt.Println("无法连接到 RCON 服务器:", err)
		return
	}
	defer conn.Close()

	// 发送命令并获取响应
	resp, err := conn.Execute("/say 123")
	if err != nil {
		fmt.Println("发送命令时出错:", err)
		return
	}

	fmt.Println("RCON 响应:", resp)
}
