package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sort"
	"encoding/json"
)

/*
1 读取IP列表
2 对IP进行去重，并统计每一个IP的出现次数
3 对每一个IP的归属地进行检查
4 判断归属地是否属于中国，如果不是则打印
5 输出非中国的IP清单，包含IP，出现次数，国家
*/

//读取IP列表
func LoadIPList(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("File is broken")
	}
	defer file.Close()
	data := make([]string, 0)
	reader := bufio.NewReader(file)
	for {
		linestr, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		linestr = strings.TrimSpace(linestr)
		if linestr == "" {
			continue
		}
		
		data = append(data, linestr)
	}
	return data
}

//先将数组中的值填入map，再用ok pattern判断值是否重复，重复便+1，不重复就将值的映射值改为1
func RemoveDupAndCount(iplist []string)map[string]int{
	m:=make(map[string]int)
	for i:=0;i<len(iplist);i++{
		if value,ok := m[iplist[i]];ok{
			m[iplist[i]]=value+1
		}else {
			m[iplist[i]]=1
		}
	}
	return m
}

//将IP个数从大到小排序
func Seq(ipListAfterReDump map[string]int)[]pair {

	for v,k := range ipListAfterReDump{
		p=append(p,pair{v,k})
	}
	sort.Slice(p,func(i,j int)bool{return p[i].key>p[j].key})
	return p
}


type pair struct {
	value string
	key int
}
var p []pair


type IPinfo struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data struct {
		IP string `json:"ip"`
		Country string `json:"country"`
		Province string `json:"province"`
		City string `json:"city"`
		County string `json:"county"`
		Region string `json:"region"`
		Isp string `json:"isp"`
	} `json:"data"`
}



func checkIPArea(ip string) []byte {
	url := "http://localhost:9090?ip="
	body := make([]byte, 0)
	response, err := http.Get(url + ip)
	if err != nil {
		fmt.Println("The site is down")
	} 
	body, _ = ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return body
}


func main() {
	filename := os.Args[1]
	ipList := LoadIPList(filename) // 读取IP列表
	ipCount := RemoveDupAndCount(ipList) //对IP进行去重，并统计每一个IP的出现次数
	ipCountSeq := Seq(ipCount)
	var j int
	var IP_Info IPinfo
	for i:=0;i<len(ipCountSeq);i++ {
		ipCheck := checkIPArea(ipCountSeq[i].value)
		err := json.Unmarshal(ipCheck, &IP_Info) //对每一个IP的归属地进行检查
		if err != nil{
		fmt.Printf("The error is %+v \n", err)
	}
	if IP_Info.Data.Country != "中国"{ //判断归属地是否属于中国，如果不是则打印
		fmt.Printf("IP: %v belongs to %v, count: %v \n",IP_Info.Data.IP,IP_Info.Data.Country,ipCountSeq[i].key) //输出非中国的IP清单，包含IP，出现次数，国家
		j++
	}
		
	}	
	fmt.Printf("非中国大陆地区的IP出现了 %v 次 \n",j)
	
}
