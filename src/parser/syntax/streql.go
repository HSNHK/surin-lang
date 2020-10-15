package syntax

import (
	"../../interpreter"
	"fmt"
	"strings"
)
//streql syntax
func StrEql(code string){
	//streql("hello"[,]"hello")
	codeSplit:=strings.Split(code,",")
	fmt.Println(interpreter.StrEql(
		//streql(["hello"],"hello")
		strings.TrimSpace(codeSplit[0][8:len(codeSplit[0])-1]),
		//streql("hello",["hello"])
		strings.TrimSpace(codeSplit[1][1:len(codeSplit[1])-2])))
}