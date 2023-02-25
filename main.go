package main

import (
	"github.com/cbot918/wsy/src/wsy"
)

func main(){
	
	port := "2346"
	
	w := wsy.New(port)
	w.Run()
	// w.ParseJson()
	// w.PrintChar(20)
}