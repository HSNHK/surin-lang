package syntax

import (
	"../../interpreter"
	"fmt"
	"strings"
)
//find function syntax
func Find(code string)  {
	//find("hello"[,]"h")
	codeSplit:=strings.Split(code,",")
	//find(["hello"],"h")
	fmt.Println(interpreter.Find(strings.TrimSpace(codeSplit[0][6:len(codeSplit[0])-1]),
		//find("hello",["h"])
		strings.TrimSpace(codeSplit[1][1:len(codeSplit[1])-2])))
}
