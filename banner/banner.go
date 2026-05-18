package main

import (
	"fmt"
	"strings"
)

func main(){
banner("Go",6)
banner("U",6)

}

func banner(text string, width int){
	fmt.Println("text", text)
	fmt.Println("repeat",strings.Repeat("\U0001F60A", width))
	fmt.Println("\U0001F60A")
	var ch rune ='A'
	fmt.Println("ch: ", ch)
	fmt.Printf("%c\n ", ch)
	fmt.Printf("%T\n ", ch)

}
