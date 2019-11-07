/*DES加密算法，为对称加密算法中的一种。是以64比特的明文为一个单位来进行加密，超过64比特的数据，要求按固定的64比特的大小分组
DES使用的密钥长度为64比特，但由于每隔7个比特设置一个奇偶校验位，因此其密钥长度实际为56比特
奇偶校验为最简单的错误检测码，即根据一组二进制代码中1的个数是奇数或偶数来检测错误
加密模式：CBC模式 全称Cipher Block Chaining模式，译为密文分组链接模式
加密步骤如下：

1）首先将数据按照8个字节一组进行分组得到D1D2......Dn（若数据不是8的整数倍，用指定的PADDING数据补位）

2）第一组数据D1与初始化向量I异或后的结果进行DES加密得到第一组密文C1（初始化向量I为全零）

3）第二组数据D2与第一组的加密结果C1异或以后的结果进行DES加密，得到第二组密文C2

4）之后的数据以此类推，得到Cn

5）按顺序连为C1C2C3......Cn即为加密结果。

当明文长度不为分组长度的整数倍时，需要在最后一个分组中填充一些数据使其凑满一个分组长度。

PKCS5Padding
加密前：数据字节长度对8取余，余数为m，若m>0,则补足8-m个字节，字节数值为8-m，即差几个字节就补几个字节，字节数值即为补充的字节数，若为0则补充8个字节的8
解密后：取最后一个字节，值为m，则从数据尾部删除m个字节，剩余数据即为加密前的原文。
加密字符串为为AAA，则补位为AAA55555;加密字符串为BBBBBB，则补位为BBBBBB22；加密字符串为CCCCCCCC，则补位为CCCCCCCC88888888。*/
package main

import (
	"fmt"
	"crypto/cipher"
	"crypto/des"
	"bytes"
	"encoding/base64"
	"bufio"
	"os"

)


func EncryptDES(plain, key[]byte){ //创建一个加密算法，需要明文和密钥

	block,_:=des.NewCipher(key) //将字符串密钥转化成为块......NewCipher creates and returns a new cipher.Block......func NewCipher(key []byte) (cipher.Block, error)
	plain = Padding(plain,block.BlockSize()) //明文补码
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

