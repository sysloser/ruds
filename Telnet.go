/* 1.输入一个域名和端口
2.将域名解析为IP
3.本地对域名和端口进行探测
4.如果5秒未获得正确返回，则打印“端口不可达”信息，并退出
5.如果获得正确返回，则打印“端口开放信息”，并退出
 */

package main
import (
	"fmt"
	"net"
	"os"
	"time"
)

func portscan(IPAddr *net.IPAddr, Port string){
	combine:= fmt.Sprintf("%v:%v",IPAddr,Port)
	//Sprintf 根据于格式说明符进行格式化并返回其结果字符串，也就是把两个值合并成1个值了。
	//如果用Printf的话，编译报错：1 variable but fmt.Printf returns 2 values，就是有2个值了
	conn,err:=net.DialTimeout("tcp",combine,5*time.Second)
	if err != nil {
		fmt.Println("端口不可达")
		//fmt.Println(conn)
		return
	}
	if conn != nil{
		fmt.Println("端口可达")
		//fmt.Println(conn)
	}
	return
}

func main(){

	host:=os.Args[1]
	port:=os.Args[2]
	addr,err := net.ResolveIPAddr("ip",host) //域名解析函数
	if err != nil {
		fmt.Printf("The domain name: %s cant'be resolved,error message is %s",host,err)
		return
	}
	fmt.Printf("The IP address is %v \n",addr)
	time.Sleep(2*time.Second)
	fmt.Printf("connecting..........%v\n",port)
	time.Sleep(2*time.Second)
	portscan(addr,port)

}
