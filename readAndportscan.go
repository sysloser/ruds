package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

/*func Isdigital(word string) (a bool) {
	pattern := "\\d+"
	result, _ := regexp.MatchString(pattern, word)
	return result
}
*/

func portscan(IPAddr *net.IPAddr, Port string) {
	combine := fmt.Sprintf("%v:%v", IPAddr, Port)
	//Sprintf 根据于格式说明符进行格式化并返回其结果字符串，也就是把两个值合并成1个值了。
	//如果用Printf的话，编译报错：1 variable but fmt.Printf returns 2 values，就是有2个值了
	conn, err := net.DialTimeout("tcp", combine, 5*time.Second)
	if err != nil {
		fmt.Println("端口不可达")
		//fmt.Println(conn)
		return
	}
	if conn != nil {
		fmt.Println("端口可达")
		//fmt.Println(conn)
	}
	return
}


func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("file is broken", err)
	}
	defer file.Close()
	//data := make([]byte, 0)
	reader := bufio.NewReader(file)
	//var combine string
	for {
		linestr, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		linestr = strings.TrimSpace(linestr)
		if linestr == "" {
			continue
		}
		for i:=0;i<len(linestr);i++{
			if string(linestr[i])==" "{
				host:=string(linestr[:i])
				addr, err := net.ResolveIPAddr("ip", host) //域名解析函数
				if err != nil {
					fmt.Printf("The domain name: %s cant'be resolved,error message is %s", host, err)
					return
				}
				port:=string(linestr[i+1:])
				fmt.Println(linestr)
				portscan(addr,port)

				//fmt.Println(combine)
				//data = append(data, combine...)
			}


		}


		//data = append(data,linestr...)

	}
	//fmt.Println(string(data))

}
