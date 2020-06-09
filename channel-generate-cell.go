package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)


func Random(s string) {
	num := [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 0; i < 8; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(9))
		s = s + num[int(n.Int64())]
	}
	fmt.Println(s)
}

func cell(c chan int){
	for{
		data:= <-c
		if data==0{
			break
		}
		Random("138")
	}
	c<-0
}

func main(){
	c:=make(chan int)
	//var b string = "138"
	//go Random(b)
	go cell(c)
	for i:=1;i<=1000;i++{
		c<-i
	}
	c<-0
	<-c

}
