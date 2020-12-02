package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

//发送文件到服务端
func SendFile(conn net.Conn) {
	f, err := os.Open("test")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	var count int64
	for {
		buf := make([]byte, 2048)
		//读取文件内容
		n, err := f.Read(buf)
		if err != nil && io.EOF == err {
			fmt.Println("文件传输完成")
			//告诉服务端结束文件接收
			conn.Write([]byte("finish"))
			return
		}
		//发送给服务端
		conn.Write(buf[:n])

		count += int64(n)
		
	}
	value := fmt.Sprintf("%.2f",count )
		//打印上传进度
		fmt.Println("文件上传：" + value + "%")
	time.Sleep(time.Second * 360)  
}

func main() {
	//创建切片，用于存储输入的路径


	//创建客户端连接
	conn, err := net.Dial("tcp", "193.167.100.100:18089")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("test"))
	buf := make([]byte, 2048)
	//读取服务端内容
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	revData := string(buf[:n])
	if revData == "ok" {
		//发送文件数据
		SendFile(conn)
	}
}
