package main

import "fmt"

func main(){
	fmt.Println(deletes("abcdea"))
	//	//fmt.Println("hello")
}
//func getLongString(s string) {
//	//arry:=make(map[string]int)
//	n := len([]rune(s))
//	//	for idex,val:=range s{
//	//        for
//	//}
//
//	for i := 0; i < n; i++ {
//		for j := i + 1; j < n; j++ {
//			if s[i] == s[j] {
//				delete()
//			}
//		}
//	}
//}
	func deletes(s string) string {
		countMap := make(map[rune]int, 0)
		//统计出现次数
		for _, v := range s {
		countMap[v]++
	}

	//	var Count int
	//	for _, v := range countMap {
	//	if Count == 0 || v >1 {
	//	Count = v
	//}
	//}
		//删除字符串中出现次数为Count的字符
		for i := len(s) - 1; i >= 0; i-- {
		if countMap[rune(s[i])] >1 {
		s = s[:i] + s[i+1:]
	}
	}
		return s
	}


