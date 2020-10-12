package syntax

import (
	"../../interpreter"
	"fmt"
	"strings"
)

func  Var(code,_type string,stack *map[string][]interface{})  {
	if _type==STRING{
		interpreter.CreateVariable(code[4:len(code)-5],"str",stack)
	}else if _type==INT{
		interpreter.CreateVariable(code[4:len(code)-5],"int",stack)
	}
}
func Push(code,_type string,stack *map[string][]interface{})  {
	if _type==STRING{
		codeSplit:=strings.Split(code,",")
		if interpreter.ExistVariable(codeSplit[0][6:len(codeSplit[0])-1],stack){
			interpreter.PushStack(codeSplit[0][6:len(codeSplit[0])-1],STRING,
				codeSplit[1][1:len(codeSplit[1])-2],stack)
		}else{
			interpreter.Log("not found variable","push str function",2)
		}
	}else if _type ==INT{
		codeSplit:=strings.Split(code,",")
		if interpreter.ExistVariable(codeSplit[0][6:len(codeSplit[0])-1],stack){
			interpreter.PushStack(codeSplit[0][6:len(codeSplit[0])-1],INT,
				codeSplit[1][:len(codeSplit[1])-1],stack)
		}else{
			interpreter.Log("not found variable","push int function",2)
		}
	}
}

func Pop(code string,stack *map[string][]interface{})  {
	if interpreter.ExistVariable(code[4:len(code)-1],stack) {
		fmt.Println(interpreter.GetValue(code[4:len(code)-1], stack))
	}else{
		interpreter.Log("not found variable","pop function",2)
	}
}

func Rm(code string,stack *map[string][]interface{})  {
	if interpreter.DeleteVariable(code[3:len(code)-1],stack)!=true{
		interpreter.Log("not found variable","rm function",2)
	}
}

func Add(code,_type string,stack *map[string][]interface{})  {
	if _type==STRING{
		codeSplit:=strings.Split(code,",")
		if interpreter.ExistVariable(codeSplit[0][5:len(codeSplit[0])-1],stack) {
			interpreter.AddToVariable(codeSplit[0][5:len(codeSplit[0])-1], "str", codeSplit[1][1:len(codeSplit[1])-2], stack)
		}else{
			interpreter.Log("not found variable","add str function",2)
		}
	}else if _type==INT{
		codeSplit:=strings.Split(code,",")
		if interpreter.ExistVariable(codeSplit[0][5:len(codeSplit[0])-1],stack) {
			interpreter.AddToVariable(codeSplit[0][5:len(codeSplit[0])-1], "int", codeSplit[1][:len(codeSplit[1])-1], stack)
		}else{
			interpreter.Log("not found variable","add int function",2)
		}
	}
}

func Stype(code string,stack *map[string][]interface{})  {
	if interpreter.ExistVariable(code[5:len(code)-1],stack){
		fmt.Println(interpreter.Type(code[5:len(code)-1],stack))
	}else{
		interpreter.Log("not found variable","add int function",2)
	}
}

func Id(code string,stack *map[string][]interface{})  {
	if interpreter.ExistVariable(code[3:len(code)-1],stack){
		fmt.Println(interpreter.Id(code[3:len(code)-1],stack))
	}else{
		interpreter.Log("not found variable","add int function",2)
	}
}