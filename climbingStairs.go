package main
import "fmt"

func clmStairs(n int)int{
	if n<=2{
		return n
	}
	result:=clmStairs(n-1)+clmStairs(n-2)
	return result
}

func main(){
	result:=clmStairs(5)
	fmt.Println(result)
}
