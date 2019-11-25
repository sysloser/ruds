package main

import (
	"bufio"
	"fmt"
	"regexp"
	"os"
	"strings"

)

func Isdigital(word string)(a bool){
	pattern := "\\d+"
	result, _ := regexp.MatchString(pattern,word)
	return result
}

func main(){
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil{
		fmt.Println("file is broken",err)
	}
	defer file.Close()
	data := make([]byte,0)
	reader := bufio.NewReader(file)
	//var section string
	for{
		linestr, err := reader.ReadString('\n')
		if err != nil{
			break
		}
		linestr = strings.TrimSpace(linestr)
		if linestr == ""{
			continue
		}
		//length:=len(linestr)
		if Isdigital(linestr[:10]) == true || linestr[:5]=="merge"{
			data = append(data,linestr...)
		}
		//data = append(data,section...)
		//break
	}
	fmt.Println(string(data))
}
