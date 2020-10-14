package syntax

import (
	"../../interpreter"
	"strconv"
	"strings"
)

func  Ivar(code string,stack *map[string][]interface{})  interface{}{
	if interpreter.ExistVariable(code[4:len(code)-1],stack) {
		return interpreter.GetValue(code[4:len(code)-1], stack)
	}else{
		interpreter.Log("not found variable","pop function",2)
		return nil
	}
}

func IF(code string)  {
	codeSpit :=strings.Split(code,",")
	v1,err1:=strconv.Atoi(codeSpit[0][3:])
	trim:=strings.Split(codeSpit[2],")")
	v2,err2:=strconv.Atoi(trim[0])
	if err1==nil && err2==nil{
		if interpreter.Logic(codeSpit[1],v1,v2)==true{
			codeSpit=strings.Split(code,"?")
			codeSpit=strings.Split(codeSpit[1],":")
			Print(codeSpit[0],STRING)
		}else {
			codeSpit=strings.Split(code,"?")
			codeSpit=strings.Split(codeSpit[1],":")
			Print(codeSpit[1],STRING)
		}
	}else{
		interpreter.Log("if format error | simple :  if(5,>=,3)?print(\"ok\"):print(\"no\")","if Syntax functions",2)
	}
}
