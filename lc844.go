//backspace  string compare

package main
import "fmt"

func backspacecompare(s string ,t string) bool{
	return build(s)==build(t)
}

func build(str string) string{
stack :=[]rune{}
for _, ch:=range str{
	if ch != '#'{
		stack= append(stack,ch)
	}else if len(stack)>0{
		stack=stack[:len(stack)-1]
	}
}
return string (stack)
}
func main(){
	s:="ab#c"
	t:="ad#c"
	fmt.Println(backspacecompare(s,t))
}