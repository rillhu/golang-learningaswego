package main

import (
	"fmt"
	"net"
	"bytes"
	"io"
)

func main() {


	fmt.Println("Hello")

	conn, err := net.Dial("tcp", "192.168.56.1:8080") //network debugger tcp server address

	fmt.Println(conn,err)


	_,err = conn.Write([]byte("HEAD / HTTP/1.0\r\n"))

	if err!=nil{
		fmt.Println("error")
	}

	var buf[512] byte
	result := bytes.NewBuffer(nil)

	for {
		n, err2 := conn.Read(buf[0:])

		fmt.Println("read done")
		if err2 != nil {
			fmt.Println("Error2:", err2)
			if err2 == io.EOF {
				break
			}
		}

		result.Write(buf[0:n])
		fmt.Println(string(result.Bytes()))
	}



	conn.Close()



}
