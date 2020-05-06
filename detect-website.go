package main

/*
1. 将需要检测的域名文件写入一个文本文件，每个URL一行
2. 对每一个URL进行探测
3. 如果探测到的URL包含“美团”或者“meituan"关键字，则打印”发现钓鱼网站：URL“
4. 如果不包含，则打印”未发现钓鱼网站“
*/

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)


func Readfile(filename string) []string {
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

func IsKeyword1(word string) (b bool) {
	pattern := "meituan"
	result, _ := regexp.MatchString(pattern, word)
	return result
}

func IsKeyword2(word string) (b bool) {
	pattern := "美团"
	result, _ := regexp.MatchString(pattern, word)
	return result
}

func main() {
	filename := os.Args[1]
	url := Readfile(filename)
	for _, k := range url {
		response,err:=http.Get(k)
		if err != nil{
			fmt.Printf("The site cannot be opened: %s \n",k)
		}else{
			body, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		if IsKeyword1(string(body)) || IsKeyword2(string(body)) {
			fmt.Printf("发现钓鱼网站：%s \n", k)
		}
		}
		
	}
}
