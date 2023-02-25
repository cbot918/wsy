package main

import (
	"fmt"
	"net"

	u "github.com/cbot918/liby/util"
)


func main(){
	fmt.Println("CLIENT")

	conn, err := net.Dial("tcp", "localhost:2000")
	u.Checke(err, "net dial failed")
	defer conn.Close()

	// fmt.Fprintf(conn, "解除封印!")
	message := "hihi"
	wn, err := conn.Write([]byte(message))
	u.Checke(err, "conn write failed")
	fmt.Println(wn)
	
	res := make([]byte, 64)
	rn, err := conn.Read(res)
	u.Checke(err, "conn read failed")
	fmt.Println(rn)

	fmt.Println(string(res))
}