package main

import "fmt"

func main(){
	/*var balance[5]float32

	a:=[3]int{1,2,3}//声明并初始化
	b:=[...]int {1,2,3,4,5,6,7}
	A:=[2][4]int{{1,2,3,4},{1,2,3, 4}}
	for index,value:=range month{

	}*/



		b := make([]int, 0, 1)
		fmt.Printf("len(b)=%d cap(b)=%d b=%v &b=%p \n", len(b), cap(b), b, b) //为什么这里没有&b[0]
		b = append(b, 1)
		fmt.Printf("len(b)=%d cap(b)=%d b=%v &b=%p &b[0]=%v\n", len(b), cap(b), b, b, &b[0])
		b = append(b, 2)
		fmt.Printf("len(b)=%d cap(b)=%d b=%v &b=%p &b[0]=%v\n", len(b), cap(b), b, b, &b[0])
		b = append(b, 3)
		fmt.Printf("len(c)=%d cap(c)=%d b=%v &b=%p &b[0]=%v\n", len(b), cap(b), b, b, &b[0])

}