package syntax

import (
	"../../interpreter"
	"fmt"
	"strconv"
	"strings"
)

func Logic(code string)  {
	codeSpit :=strings.Split(code,",")
	v1,err1:=strconv.Atoi(codeSpit[0][6:])
	v2,err2:=strconv.Atoi(codeSpit[2][:len(codeSpit[2])-1])
	if err1==nil && err2==nil{
		fmt.Println(interpreter.Logic(codeSpit[1],v1,v2))
	}else{
		interpreter.Log("logic format error simple :  logic(5,>=,3)","Logic Syntax functions",2)
	}
}
