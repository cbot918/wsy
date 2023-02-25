package wsy

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net"
	"os"
	"strings"

	u "github.com/cbot918/liby/util"
)

type Wsy struct {
	listener 						net.Listener
	FirstRequest 				string
	Data 								Data
	CurrentFrame			Frame
}

type Data struct{
	Host 										[]string
	Connection 							[]string
	Pragma 									[]string
	CacheControl						[]string
	UserAgent								[]string
	Upgrade									[]string
	Origin									[]string
	SecWebSocketVersion			[]string
	AcceptEncoding					[]string
	AcceptLanguage					[]string
	SecWebSocketKey					[]string
	SecWebSocketExtensions	[]string
	SecWebSocketAccept				string
}

type Frame struct {
	Fin						byte
	Opcode				byte
	IsMasked			byte
	PayloadLen		byte
	Mask					[]byte
	Payload				[]byte
}

func New() *Wsy{
	w := new(Wsy)
	l, err := net.Listen("tcp", "localhost:2346")
	u.Checke(err, "net listen failed")
	fmt.Println("welcome to wsy")
	w.listener = l
	return w
}

func (w *Wsy) Listen() string {

	for {
		conn, err := w.listener.Accept()
		u.Checke(err, "listener accept error")

		// message := make(chan string)
		go func (c net.Conn){
			defer conn.Close()
			req := make([]byte, 4096)
			_, err := conn.Read(req)
			u.Checke(err, "conn read failed")
			// message <- string(req)
			w.FirstRequest = string(req)
			w.ParseJson()
			fmt.Println("Incoming Request")
			fmt.Println("Websocket handshake detected with key: ", strings.TrimSpace(strings.Join(w.Data.SecWebSocketKey,"")))
			fmt.Println("Responding to handshake with key: ", w.GetReturnSec())
			response := fmt.Sprintf("HTTP/1.1 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: %s\r\n\r\n",w.GetReturnSec())
			fmt.Println("")
			res := []byte(response)
			_, err = conn.Write(res)
			u.Checke(err, "conn write failed")

			_, err = conn.Read(req)
			u.Checke(err, "conn read failed")
			// fmt.Println(string(req))
			
			fmt.Println( w.GetDecMessage(req) )
			

		}(conn)
		// w.FirstRequest = <- message
		
		// w.GetReturnSec()
		
		

	}
}

func (w *Wsy) GetDecMessage(data []byte) string {
	firstByte := data[0]
	secondByte := data[1]

	fin := firstByte & 0b10000000
	opcode := firstByte & 0b00001111
	is_masked := secondByte & 0b10000000
	payload_len := secondByte & 0b01111111
	
	fmt.Printf("fin: %d\nopcode: %d\nis_masked: %d\npayload_len: %d\n",fin,opcode,is_masked,payload_len)

	// process mask
	mask := []byte{data[2],data[3],data[4],data[5]}
	fmt.Println("mask: ",mask)

	// process payload data
	payload := []byte{}
	for i:=6; i<=int(payload_len+6); i++ {
		payload = append(payload, data[i])
	}
	fmt.Println("payload: ",payload)

	// XOR payload and mask
	result := []byte{}
	for i,item := range payload {
		result = append(result, item ^ mask[i%4])
	}
	// fmt.Println(string(result))

	return string(result)
	// for i:=0; i<10; i++  {
	// 	fmt.Println(data[i])
	// }
}

func (w *Wsy) GetReturnSec() string {
	secWebSocketKey := strings.TrimSpace(strings.Join(w.Data.SecWebSocketKey, ""))
	var keyGUID = []byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11")
	h := sha1.New()
	h.Write([]byte(secWebSocketKey))
	h.Write(keyGUID)
	secWebSocketAccept := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return secWebSocketAccept
}

func (w *Wsy) ParseJson(){

	dataArray := strings.Split(strings.ReplaceAll(w.FirstRequest, "\r\n","\n"),"\n")
	
	for index, item := range dataArray {
		if index == 0 { continue }
		key, value := w.LineToKV(item)
		switch strings.Join(key, "") {
		case "Host": { w.Data.Host = value }
		case "Connection" : { w.Data.Connection = value }
		case "Sec-WebSocket-Key" : { w.Data.SecWebSocketKey = value }
		}
	}
}

func (w *Wsy) LineToKV(line string) ([]string,[]string){
	return w.ReadBefore(":",line),w.ReadAfter(":",line)
}

func (w *Wsy) ReadAfter(targetChar string, content string) []string{
	var buf []string
	var flag bool
	for _, char := range content {
		if string(char) == targetChar { flag = true; continue }
		if (flag){ buf = append(buf, string(char))}
	}
	return buf
}


func (w *Wsy) ReadBefore(targetChar string, content string) []string{
	var buf []string
	for _, char := range content {
		if string(char) != targetChar {
			buf = append(buf, string(char))	
		} else { break }
	}
	return buf
}

func (w *Wsy) PrintChar(count int){
	i := 0
	for _, ch := range w.FirstRequest {
		i ++
		fmt.Printf("%q\n",string(ch))
		if i >= count{
			os.Exit(1)
		}
	}
}