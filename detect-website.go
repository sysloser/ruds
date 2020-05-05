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

func detect(url string) string {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return string(body)
}

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
	pattern := "baidu"
	result, _ := regexp.MatchString(pattern, word)
	return result
}

func IsKeyword2(word string) (b bool) {
	pattern := "百度"
	result, _ := regexp.MatchString(pattern, word)
	return result
}

func main() {
	filename := os.Args[1]
	url := Readfile(filename)
	for _, k := range url {
		site := detect(k)
		if IsKeyword1(site) || IsKeyword2(site) {
			fmt.Printf("发现钓鱼网站：%s \n", k)
			continue
		}
	}
}
