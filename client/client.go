/*
A very simple TCP client written in Go.

This is a toy project that I used to learn the fundamentals of writing
Go code and doing some really basic network stuff.

Maybe it will be fun for you to read. It's not meant to be
particularly idiomatic, or well-written for that matter.
*/
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)



func main() {
	conn ,err := net.Dial("udp","193.167.100.100:8089")
	if err != nil{
		fmt.Println("net.Dial err",err)
		return
	}
	defer conn.Close()

	//5.发送文件名给接收端

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("conn.Read() err:%v\n", err)
		return
	}
	fileName := string(buf[:n])

	//回写ok给发送端
	_, _ = conn.Write([]byte("ok"))


	recivefil(conn,fileName)

}
func recivefil(conn net.Conn,fileName string) {

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("os.Create() err:%v\n", err)
		return
	}
	defer file.Close()

	//从网络中读数据，写入本地文件
	for {
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)

		//写入本地文件，读多少，写多少
		file.Write(buf[:n])
		if err != nil {
			if err == io.EOF {
				fmt.Printf("recieving finish\n")
				fi,err:=os.Stat(fileName)
				if err ==nil {
					fmt.Println("file size is ",fi.Size(),"Bytes")
				}
			} else {
				fmt.Printf("conn.Read() err:%v\n", err)
			}
			return
		}

	}

}
