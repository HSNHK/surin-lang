package syntax

import (
	"../../interpreter"
	"fmt"
	"strconv"
)
//logic function syntax
func Logic(code string)  {
	NumberA, Operation,NumberB:=Three(code)
	v1,err1:=strconv.Atoi(NumberA)
	v2,err2:=strconv.Atoi(NumberB)
	if err1==nil && err2==nil{
		fmt.Println(interpreter.Logic(Operation,v1,v2))
	}else{
		interpreter.Log("logic format error simple :  logic(5,>=,3)","Logic Syntax functions",2)
	}
}
