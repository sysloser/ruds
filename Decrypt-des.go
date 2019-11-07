package main

import (
	"fmt"
	"crypto/cipher"
	"crypto/des"
	//"bytes"
	"encoding/base64"
	"bufio"
	"os"

)


func Unpad(origData []byte)[]byte{ //创建一个块还原函数
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}



func DecryptDES(secret string, key []byte){ //创建一个解密函数
	crypted, _ :=base64.StdEncoding.DecodeString(secret)
	block, _ := des.NewCipher(key)
	blockMode := cipher.NewCBCDecrypter(block,key)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData,crypted)
	origData = Unpad(origData)
	fmt.Println(string(origData))
}


func main(){
	key:=[]byte("12345678")//定义密钥
	// 准备从标准输入读取数据。
    inputReader := bufio.NewReader(os.Stdin)
    fmt.Println("请输入你需要解密的话:\n")
    // 读取数据直到碰到 \n 为止。
    input, err := inputReader.ReadString('\n')
    if err != nil {
        fmt.Printf("An error occurred: %s\n", err)
        // 异常退出。
        os.Exit(1)
    } else {
        // 用切片操作删除最后的 \n 。
        cryptedtext:=input[:len(input)-2]
        fmt.Println("你的话被解密了:\n")
        DecryptDES(cryptedtext,key)//调用解密函数
    }

	

}
/*

func EncryptDES(plain, key[]byte){ //创建一个加密算法，需要明文和密钥

	block,_:=des.NewCipher(key) //将字符串密钥转化成为块......NewCipher creates and returns a new cipher.Block......func NewCipher(key []byte) (cipher.Block, error)
	plain = Padding(plain,block.BlockSize()) //明文补码 ---BlockSize是des中的常量，8字节
	blockmode:= cipher.NewCBCEncrypter(block,key) //设置加密方式
	crypted:=make([]byte,len(plain)) //创建密文切片
	blockmode.CryptBlocks(crypted,plain) //加密明文,加密后的数据放到数组中
	fmt.Println(base64.StdEncoding.EncodeToString(crypted)) //将字节数组转换成字符串




}

func Padding(data []byte, blocksize int)[]byte{ //计算补位
	pad:=blocksize-len(data)%blocksize //计算要补多少位
	 //Repeat returns a new byte slice consisting of count copies of b. func Repeat(b []byte, count int) []byte
	 //Repeat()函数的功能是把参数一 切片复制 参数二count个,然后合成一个新的字节切片返回
 // 需要补padding位的padding值
	padtext := bytes.Repeat([]byte{byte(pad)}, pad)
	return append(data,padtext...) //补充的内容拼接回去



}

func main(){
	key:=[]byte("12345678")//创建密钥
	// 准备从标准输入读取数据。
    inputReader := bufio.NewReader(os.Stdin)
    fmt.Println("请输入你需要加密的话:\n")
    // 读取数据直到碰到 \n 为止。
    input, err := inputReader.ReadString('\n')
    if err != nil {
        fmt.Printf("An error occurred: %s\n", err)
        // 异常退出。
        os.Exit(1)
    } else {
        // 用切片操作删除最后的 \n 。
        plaintext:=[]byte(input[:len(input)-2])
        fmt.Println("你的话被加密了:\n")
        EncryptDES(plaintext,key)
    }


	//plaintext:=[]byte("Hello") //输入明文
	//EncryptDES(plaintext,key) //输出密文

}
*/
