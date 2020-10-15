package syntax

import (
	"../../interpreter"
	"fmt"
	"strings"
)

//value type
const STRING ="str"
const INT ="int"

//print syntax
func Print(code,_type string){
	if _type==STRING{
		//print string value
		interpreter.Print(code[7:len(code)-2], "str")
	}else if _type==INT{
		//print int value
		interpreter.Print(code[7:len(code)-2], "int")
	}
}
//len syntax
func Len(code string)  {
	//len(["hello"])
	fmt.Println(interpreter.DataLen(strings.TrimSpace(code[5:len(code)-2])))
}