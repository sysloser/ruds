package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)



//读取文件并输出IP列表
func OutputIPlist(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("File is broken")
	}
	defer file.Close()
	data := make([]string, 0)
	patternIP := regexp.MustCompile("((2[0-4]\\d|25[0-5]|[01]?\\d\\d?)\\.){3}(2[0-4]\\d|25[0-5]|[01]?\\d\\d?)")
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
		isIP := patternIP.FindString(linestr)

		data = append(data, isIP)
	}
	return data
}

//检查IP归属信息
func checkIParea(ip string) string {
	url := "http://localhost:9090?ip="
	body := make([]byte, 0)
	response, err := http.Get(url + ip)
	if err != nil {
		fmt.Println("The site is down")
	} else {
		body, _ = ioutil.ReadAll(response.Body)
	}
	return string(body)
}

func main() {
	filename := os.Args[1]
	ipList := OutputIPlist(filename)
	for _, k := range ipList {
		ipArea := checkIParea(k)
		fmt.Println(ipArea)
	}

}
