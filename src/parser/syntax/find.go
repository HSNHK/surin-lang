package syntax

import (
	"../../interpreter"
	"fmt"
	"strings"
)

func Find(code string)  {
	codeSplit:=strings.Split(code,",")
	fmt.Println(interpreter.Find(strings.TrimSpace(codeSplit[0][6:len(codeSplit[0])-1]),
		strings.TrimSpace(codeSplit[1][1:len(codeSplit[1])-2])))
}
