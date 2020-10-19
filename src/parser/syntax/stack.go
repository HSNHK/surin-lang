package syntax

import (
	"../../interpreter"
	"fmt"
	"strings"
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
//cmp two variable
func Cmp(code string,stack *map[string][]interface{})  {
	codesplit:=strings.Split(code,",")
	if interpreter.ExistVariable(codesplit[0][4:len(codesplit[0])],stack){
		if interpreter.ExistVariable(codesplit[1][:len(codesplit[1])-1],stack){
			fmt.Println(interpreter.Cmp(codesplit[0][4:],codesplit[1][:len(codesplit[1])-1],stack))
		}else{
			interpreter.Log("not found variable : "+codesplit[1][:len(codesplit[1])-1],"add int function",2)
		}
	}else {
		interpreter.Log("not found variable : "+codesplit[0][4:],"add int function",2)
	}
}