package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main(){
	//输入一个文本文件名，如果文件名不符合要求，则进行判断，并退出程序
	filename:=os.Args[1]
	if len(os.Args)==1 || filename=="-h" {
		fmt.Println("请输入一个文件名")
		os.Exit(1)
	}

	data:=Readfile(filename) //读取文件，将其中所有行提出来，形成string
	a:=ProcString(data) //处理string，变小写，切除前后空格和标点符号，并返回一个string数组
	d:=WordCount(a) //将统计结果传入一个map
	//打印带有值的map
	/*for v,k:=range d{
		fmt.Println(v,k)
	}*/
	//根据Key进行map的排序
	type pair struct {
		value string
		key int
	}
	var p []pair
	for v,k := range d{
		p=append(p,pair{v,k})
	}
	sort.Slice(p,func(i,j int)bool{return p[i].key>p[j].key}) //从大到小排序
	fmt.Println(p)
}

//读取文件，并将文件逐行取出，形成一个string
func Readfile(s string)string{
	file, err := os.Open(s)
	if err != nil{
		fmt.Println("file is broken",err)
	}
	defer file.Close()
	data := make([]byte,0)
	reader := bufio.NewReader(file)
	for{
		linestr, err := reader.ReadString('\n')
		if err != nil{
			break
		}
		linestr = strings.TrimSpace(linestr)
		if linestr == ""{
			continue
		}
		data=append(data,linestr...)
	}
	return string(data)
}

//将从文件中取出的string进行处理
func ProcString(words string)[]string{
	//将所有字母都变成小写
	var a string = strings.ToLower(string(words))
	//切割掉string前后的空格
	var b string = strings.TrimSpace(a)
	//切割掉string里所有标点符号
	c:=SplitNoLetter(b)
	//将string放入计数函数
	return c
}
//切割掉string中所有标点符号的函数，并返回一个string数组
func SplitNoLetter(word string)[]string{
	f := func(word rune) bool {
		return !unicode.IsLetter(word)
	}
	return strings.FieldsFunc(word, f)
}
//先将数组中的值填入map，再用ok pattern判断值是否重复，重复便+1，不重复就将值的映射值改为1
func WordCount(words []string)map[string]int{
	m:=make(map[string]int)
	for i:=0;i<len(words);i++{
		if value,ok := m[words[i]];ok{
			m[words[i]]=value+1
		}else {
			m[words[i]]=1
		}
	}
	return m
}
