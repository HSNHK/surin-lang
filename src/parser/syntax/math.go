package syntax

import (
	"../../interpreter"
	"fmt"
	"strconv"
)
//math function syntax
func Math(code string)  {
	NumberA, Operation,NumberB:=Three(code)
	v1,err1:=strconv.Atoi(NumberA)
	v2,err2:=strconv.Atoi(NumberB)
	if err1==nil && err2==nil{
		fmt.Println(interpreter.Math(Operation,v1,v2))
	}else{
		interpreter.Log("logic format error simple : math(5,*,2)","Math Syntax function",2)
	}
}