package syntax

import (
	"../../interpreter"
	"fmt"
)

//get variable type syntax
func Stype(code string,stack *map[string][]interface{})  {
	//check exist variable
	if interpreter.ExistVariable(code[5:len(code)-1],stack){
		//type([variable-name])
		fmt.Println(interpreter.Type(code[5:len(code)-1],stack))
	}else{
		interpreter.Log("not found variable","add int function",2)
	}
}
//get variable id syntax
func Id(code string,stack *map[string][]interface{})  {
	//check exist variable
	if interpreter.ExistVariable(code[3:len(code)-1],stack){
		//id([variable-name])
		fmt.Println(interpreter.Id(code[3:len(code)-1],stack))
	}else{
		interpreter.Log("not found variable","add int function",2)
	}
}