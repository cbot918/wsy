package wsy

import (
	"fmt"
	"net"

	u "github.com/cbot918/liby/util"
)

type Wsy struct {
	listener 						net.Listener
}

func New() *Wsy{
	w := new(Wsy)
	l, err := net.Listen("tcp", "localhost:2346")
	u.Checke(err, "net listen failed")
	fmt.Println("[*] welcome to wsy")
	w.listener = l
	return w
}

func (w *Wsy) Run() string {

	for {
		conn, err := w.listener.Accept()
		u.Checke(err, "listener accept error")

		go func (c net.Conn){
			defer conn.Close()

			ch := NewConnHandler(conn)	
			
			req := ch.ReadSocket()												// 讀取web ws發過來的第一個http request
			
			err := ch.Upgrade(string(req))								// 將連線升級成 websocket
			u.Checke(err, "upgrade write socket failed")
			fmt.Println("[*] ws upgrade success!")

			res := ch.ReadSocket()												// 寫死: 預設web會發一個message過來
			fmt.Println("[*] receved message")

			message := ch.DecodeFrame(res)								// 根據Spec解碼Frame把message取出來
			fmt.Println("message: ", string(message))

		}(conn)
	}
}









// func (w *Wsy) GetMessage() string {
	
// }


// func (c *ConnHandler) PrintChar(count int){
// 	i := 0
// 	for _, ch := range w.FirstRequest {
// 		i ++
// 		fmt.Printf("%q\n",string(ch))
// 		if i >= count{
// 			os.Exit(1)
// 		}
// 	}
// }