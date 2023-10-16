package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	println("Ожидание подключения")
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	println("Connected")
	defer conn.Close()

	for {
		resp := NewResp(conn)
		val, err := resp.Read()
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println(val)
		//data := make([]byte, 1024)
		//_, err = conn.Read(data)
		//if err != nil {
		//	if err == io.EOF {
		//		fmt.Println("Bye-bye user")
		//		break
		//	}
		//	fmt.Println("Cant user input")
		//}
		conn.Write([]byte("+ShutUp\r\n"))
	}
}
