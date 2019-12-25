//中年人的第一款端口扫描器
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)
var input *string

func init(){
	input=flag.String("n","","Please input IP address or domain name")
	if flag.NArg() == 0{
		fmt.Println("help message")
		os.Exit(1)
	}
}

func portscan(IPAddr *net.IPAddr) []int {
	var PortOpen = make([]int, 0)
	var tcpport int
	for tcpport = 1; tcpport <= 65535; tcpport++ {
		combine := fmt.Sprintf("%v:%v", IPAddr, tcpport)
		conn, err := net.DialTimeout("tcp", combine, 1*time.Second)
		if err != nil {
			fmt.Printf("The port %d ......closed\n",tcpport)
		}
		if conn != nil {
			fmt.Printf("The port %d .....opened\n",tcpport)
			PortOpen = append(PortOpen, tcpport)
		}
	}
	return PortOpen
}

func main(){
	flag.Parse()
	host,err:=net.ResolveIPAddr("ip",string(*input))
	if err != nil{
		fmt.Println(err)
	}
	result:=portscan(host)
	fmt.Println(*input)
	fmt.Println("Opened ports:",result)
}
