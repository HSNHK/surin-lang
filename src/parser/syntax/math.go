package syntax

import (
	"../../interpreter"
	"fmt"
	"strconv"
	"strings"
)

func Math(code string)  {
	codeSpit :=strings.Split(code,",")
	v1,err1:=strconv.Atoi(codeSpit[0][5:])
	v2,err2:=strconv.Atoi(codeSpit[2][:len(codeSpit[2])-1])
	if err1==nil && err2==nil{
		fmt.Println(interpreter.Math(codeSpit[1],v1,v2))
	}else{
		interpreter.Log("logic format error simple : math(5,*,2)","Math Syntax function",2)
	}
}