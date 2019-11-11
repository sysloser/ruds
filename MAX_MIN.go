/*最小值和最大值
编写一个函数，找到 int slice ( []int ) 中的最大值。
编写一个函数，找到 int slice ( []int ) 中的最小值。*/
package main
import "fmt"

func main(){
	a:=[]int{-1,-1999,2,3,43,2,6,5,0,2225,323,5}
	MAX_ARRY(a)
	MIN_ARRY(a)
}

func MAX_ARRY(numbers []int){
	length:=len(numbers)-1
	for i:=0;i<=length;i++{
		for j:=i;j<=length;j++{
			if numbers[j]<numbers[i]{
				t:=numbers[j]
				numbers[j] = numbers[i]
				numbers[i] = t
			}	
		}
	}
	fmt.Println("The Max_number is",numbers[length])

}

func MIN_ARRY(numbers []int){
	length:=len(numbers)-1
	for i:=0;i<=length;i++{
		for j:=i;j<=length;j++{
			if numbers[j]<numbers[i]{
				t:=numbers[j]
				numbers[j] = numbers[i]
				numbers[i] = t
			}	
		}
	}
	fmt.Println("The Min_number is",numbers[0])

}
