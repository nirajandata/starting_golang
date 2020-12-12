package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	const l_num = 30
	loader:= fmt.Sprintf("\x0c[%%-%vs]", l_num) // {cls ra %-30 auxa result } + garda chai opposite ma move hunxa 
	for i:=0; i <l_num; i++ {
	 fmt.Printf(loader, strings.Repeat("=", i)+">")
	 time.Sleep(200 * time.Millisecond)
	}
	fmt.Printf(loader+" Uff!", strings.Repeat("=", l_num))
}
