package syntax

import (
	"../../interpreter"
	"fmt"
	"strings"
)

func StrEql(code string){
	codeSplit:=strings.Split(code,",")
	fmt.Println(interpreter.StrEql(
		strings.TrimSpace(codeSplit[0][8:len(codeSplit[0])-1]),
		strings.TrimSpace(codeSplit[1][1:len(codeSplit[1])-2])))
}