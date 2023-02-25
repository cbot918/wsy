package main

import (
	"github.com/cbot918/wsy/src/wsy"
)

func main(){
	w := wsy.New()
	w.Listen()
	w.ParseJson()
	// w.PrintChar(20)
}