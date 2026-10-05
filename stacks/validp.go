package main
import "fmt"
func validParen(s string) bool{
	stack:=[]rune{}

	for _, ch:=range s{
		if ch=='{' || ch=='[' || ch=='('{
			stack=append(stack,ch)
			continue
		}
		if len(stack)==0{
			return false
		}
		top:=stack[len(stack)-1]

		if top!='{' && ch=='}' ||
		   top!='[' && ch==']' ||
		   top!='(' && ch==')' {
		   return false
			}
		stack=stack[:len(stack)-1]
	}
	return len(stack)==0

}

func main(){
	s:="({[]})"
fmt.Println(validParen(s))
}