package register_syntax

import (
	"../../../interpreter"
	"../../../interpreter/register-structure"
	"fmt"
	"strings"
)
//value type
const STRING ="str"
const INT ="int"

//get variable type syntax
func Stype(code string, register *map[string][]interface{})  {
	//check exist variable
	if register_structure.ExistVariable(code[5:len(code)-1], register){
		//type([variable-name])
		fmt.Println(register_structure.Type(code[5:len(code)-1], register))
	}else{
		interpreter.Log("not found variable","add int function",2)
	}
}
//get variable id syntax
func Id(code string, register *map[string][]interface{})  {
	//check exist variable
	if register_structure.ExistVariable(code[3:len(code)-1], register){
		//id([variable-name])
		fmt.Println(register_structure.Id(code[3:len(code)-1], register))
	}else{
		interpreter.Log("not found variable","add int function",2)
	}
}
//cmp two variable
func Cmp(code string, register *map[string][]interface{})  {
	codesplit:=strings.Split(code,",")
	if register_structure.ExistVariable(codesplit[0][4:len(codesplit[0])], register){
		if register_structure.ExistVariable(codesplit[1][:len(codesplit[1])-1], register){
			fmt.Println(register_structure.Cmp(codesplit[0][4:],codesplit[1][:len(codesplit[1])-1], register))
		}else{
			interpreter.Log("not found variable : "+codesplit[1][:len(codesplit[1])-1],"add int function",2)
		}
	}else {
		interpreter.Log("not found variable : "+codesplit[0][4:],"add int function",2)
	}
}