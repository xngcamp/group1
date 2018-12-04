package main
func fobon(){
	//var f[]int
	f:=[9]int{1,1}
	for i:=2;i<9;i++{
		f[i]=f[i-1]+f[i-2]
	}
}
