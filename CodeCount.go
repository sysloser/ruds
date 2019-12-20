//计算一个文件里有多少行代码，多少注释，多少行空格
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	filename := os.Args[1]
	if len(os.Args)<2||os.Args[1]=="-h"{
		fmt.Println("Please input a filename")
		os.Exit(0)
	}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("file is broken", err)
		os.Exit(0)
	}
	data := make([]byte, 0)
	defer file.Close()
	notecount:=make([]int,0)
	reader := bufio.NewReader(file)
	var  count,note1,noteslash,BlankCount,combinecount,combinecount2 int

	for {
		linestr, err := reader.ReadString('\n') //读取换行符前的所有内容
		if err != nil {
			break
		}
		count++
		if SearchNote3(linestr){
			noteslash++
		}
		if len(linestr)-1==CountBlank(linestr){ //计算空白行
			BlankCount++
		}

		if SearchNote1(linestr) {
			note1 = count
		}
		if SearchNote2(linestr) {
			note2:=count
			c:=note2 - note1+1
			notecount=append(notecount,c)
		}
		if SearchNote3(linestr)&&string(linestr[0])!="/"{
			combinecount++
		}
		if SearchNote1(linestr)&&string(linestr[0])!="/"{
			combinecount2++
		}
		data = append(data,fmt.Sprintf("%v - %v",count,linestr)...)
	}
	var d int
	for i:=0;i<len(notecount);i++{
		d=d+notecount[i]
	}
	codecount:=count-BlankCount-(d+noteslash)+combinecount+combinecount2
	fmt.Printf("文件总共有%v行\n",count)
	fmt.Printf("其中空行总共有%v行\n",BlankCount)
	fmt.Printf("注释总共有%v行\n",d+noteslash)
	fmt.Printf("又是双斜杠注释又是代码的有%v行\n",combinecount)
	fmt.Printf("又是斜杠星号又是代码的有%v行\n",combinecount2)
	fmt.Printf("代码总共有%v行\n",codecount)
	fmt.Printf(string(data))
}

func SearchNote1(s string)bool{  //双引号转一次，正则转一次
	pattern := "\\/\\*"
	result,_:=regexp.MatchString(pattern,s)
	return result
}
func SearchNote2(s string)bool{
	pattern := "\\*\\/"
	result,_:=regexp.MatchString(pattern,s)
	return result
}
func SearchNote3(s string)bool{
	pattern := "//"
	result,_:=regexp.MatchString(pattern,s)
	return result
}
func CountBlank(s string)int{
	rBlank:=regexp.MustCompile(" ")
	Blankcount:=len(rBlank.FindAllStringSubmatch(s,-1))
	return Blankcount
}

