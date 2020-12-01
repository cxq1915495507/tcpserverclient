/*
A very simple TCP server written in Go.

This is a toy project that I used to learn the fundamentals of writing
Go code and doing some really basic network stuff.

Maybe it will be fun for you to read. It's not meant to be
particularly idiomatic, or well-written for that matter.
*/
package main

import (
	_ "bufio"
	"fmt"
	"io"
	"net"
	"os"
	_ "strconv"
	_ "time"
)


func main() {
	listener ,err := net.Listen("tcp","127.0.0.1:18089")
	if err != nil{
		fmt.Println("net.Listener err:",err)
		return
	}
	defer listener.Close()//关闭socket


	//2.阻塞监听客户端连接请求,成功建立连接，返回用于通信的socket---conn
	conn ,err := listener.Accept()
	if err != nil{
		fmt.Println("listener.Accept err:",err)
		return
	}
	defer conn.Close()//关闭socket
	_, err = conn.Write([]byte("test"))


	//3.从conn套接字中获取文件名，写入缓存buf中
	buf := make([]byte,4096)
	n ,err := conn.Read(buf)
	if err != nil{
		fmt.Println("conn.Read err:",err)
		return
	}

	//4.从buf中提取文件名
	//5.回写给发送端ok

	if "ok" == string(buf[:n]) {
		//8.是ok，写文件内容给接收端--借助conn
		fmt.Println("ok")
		fi, err := os.Open("test")
		if err != nil {
			fmt.Println("os.Open err", err)
			return
		}
		buf := make([]byte,4096)
		for{
			n, err = fi.Read(buf)
			if err == io.EOF {
				fmt.Println("finish")
				return
			}
			_, _ = conn.Write(buf[:n])
		}

	}
	_, _ = conn.Write([]byte("ok"))

	//6.获取文件内容
	go recivefile(conn)
}

func recivefile(conn net.Conn)  {

	//6.2从网络socket中读数据，写入本地文件中
	buf := make([]byte,4096)
	var count = 0

	for  {

		n,_ := conn.Read(buf) //从conn中读数据到buf中

		if n == 0{  //判断是否读取数据完毕
			fmt.Println("receiving finish: total bytes is")
			fmt.Println(count)
			return
		}

		//将buf中的数据写入到本地文件
		count =+n
	}

}
