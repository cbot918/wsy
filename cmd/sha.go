package main

import (
	"crypto/sha1"
	"fmt"
)

func main(){

	// //这是将要编解码的字符串。
	// data := "abc123!?$*&()'-=@~"

	// sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	// fmt.Println(sEnc)
	// //解码可能会返回错误，如果不确定输入信息格式是否正确，那么，你就需要进行错误检查了。
	// sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	// fmt.Println(string(sDec))
	// fmt.Println()
	// //使用 URL 兼容的 base64 格式进行编解码。
	// uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	// fmt.Println(uEnc)
	// uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	// fmt.Println(string(uDec))
	// //标准 base64 编码和 URL 兼容 base64 编码的编码字符串存在稍许不同（后缀为 + 和 -），但是两者都可以正确解码为原始字符串。

	str := "stringgggg"
	s :=sha1.New()
	s.Write([]byte(str))
	bs := s.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x",bs)
	// base64.URLEncoding.EncodeToString([]byte(strByte))
	
}