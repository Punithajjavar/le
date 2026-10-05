package main
import (
	"sort"
	"fmt"
)
func checkIfCanBreak(s1 string,s2 string) bool {
	a:=[]byte(s1)
	b:=[]byte(s2)
	sort.Slice(a,func(i,j int) bool{
		return a[i]<a[j]
	})
	sort.Slice(b,func(i,j int)bool{
		return b[i]<b[j]
	})
	s1BreaksS2:=true
	s2BreacksS1:=true

	for i:=0;i<len(a);i++{
		if a[i]<b[i]{
			s1BreaksS2=false
		}
		if b[i]<a[i]{
			s2BreacksS1=false
		}
	}
	return s1BreaksS2 || s2BreacksS1
}

func main(){
s1:="bzc"
s2:="xaz"
result:=checkIfCanBreak(s1,s2)
fmt.Println(result)
}