package syntax

import (
	"../../interpreter"
	"fmt"
	"strconv"
	"strings"
)
//math function syntax
func Math(code string)  {
	//math(2[,]+[,]2)
	codeSpit :=strings.Split(code,",")
	//math([2],+,2)
	v1,err1:=strconv.Atoi(codeSpit[0][5:])
	//math(2,+,[2])
	v2,err2:=strconv.Atoi(codeSpit[2][:len(codeSpit[2])-1])
	if err1==nil && err2==nil{
		//math(2,[+],2)
		fmt.Println(interpreter.Math(codeSpit[1],v1,v2))
	}else{
		interpreter.Log("logic format error simple : math(5,*,2)","Math Syntax function",2)
	}
}