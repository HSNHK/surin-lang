package syntax

import (
	"../../interpreter"
	"fmt"
	"strings"
)

const STRING ="str"
const INT ="int"

func Print(code,_type string){
	if _type==STRING{
		interpreter.Print(code[7:len(code)-2], "str")
	}else if _type==INT{
		interpreter.Print(code[7:len(code)-2], "int")
	}
}

func Len(code string)  {
	fmt.Println(interpreter.DataLen(strings.TrimSpace(code[5:len(code)-2])))
}