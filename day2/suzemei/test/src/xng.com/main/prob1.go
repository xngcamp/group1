package main
import(
	"fmt"
)

func main() {
	var x int

	fmt.Scanf("%d",&x)
	var a = make([]int, x, x)
	for i := 0; i < x; i++ {
		fmt.Scanf("%d",&a[i])
	}
	//for i := 0; i < x; i ++ {
	//	fmt.Println(a[i])
	//}

	for i := 0; i < x; i++ {
		var b int = 0
		var c int = 0
		b=a[i] - i
		for j:= i + 1; j<x;j++ {
			if b==a[j] - j{
				c++
			}else{
					break
			}

		}
		if a[i] == 1 {
			fmt.Println(c)
		} else if a[x-1] == 1000&&c>0 {
			fmt.Println(c)
		} else {
			fmt.Println(c)
		}
	}

}
