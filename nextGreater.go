package main
import "fmt"
//leetcode 496
func nextGreaterElement(num1 [] int,num2 []int) []int{
	stack :=[]int{}
	nextGreater:=make(map[int]int)

	//find next greater element for num2 
for i:=len(num2)-1;i>=0;i--{
	x:=num2[i]

	for len(stack)>0 && stack[len(stack)-1]<=x{
		stack = stack[:len(stack)-1]
	}

	//find next greater 
	if len(stack)==0{
		nextGreater[x]=-1
	}else {
			nextGreater[x]=stack[len(stack)-1]
		}
// put x into stack
		stack =append(stack,x)
	}
ans := make ([]int,len(num1))

for i:=0;i<len(num1);i++{
	ans[i]=nextGreater[num1[i]]
}
return ans
	}

func main(){
	num1 :=[]int{9,10}
	num2:=[]int{2,1,9,10}
	ans:=nextGreaterElement(num1,num2)
	fmt.Println(ans)
}