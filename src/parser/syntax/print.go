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
		interpreter.Print(One(code), "str")
	}else if _type==INT{
		//print int value
		interpreter.Print(One(code), "int")
	}
}
//len syntax
func Len(code string)  {
	//len(["hello"])
	fmt.Println(interpreter.DataLen(strings.TrimSpace(One(code))))
}