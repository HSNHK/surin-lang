package Variable_syntax

import (
	"../../../interpreter"
	"../../../interpreter/register-structure"
	"../../syntax"
	"fmt"
)
//value type
const STRING ="str"
const INT ="int"

//get variable type syntax
func Stype(code string, register *map[string][]interface{})  {
	//check exist variable
	VariableName:=syntax.One(code)
	if register_structure.ExistVariable(VariableName, register){
		fmt.Println(register_structure.Type(VariableName, register))
	}else{
		interpreter.Log("not found variable","add int function",2)
	}
}
//get variable id syntax
func Id(code string, register *map[string][]interface{})  {
	//check exist variable
	VariableName:=syntax.One(code)
	if register_structure.ExistVariable(VariableName, register){
		fmt.Println(register_structure.Id(VariableName, register))
	}else{
		interpreter.Log("not found variable","add int function",2)
	}
}
//cmp two variable
func Cmp(code string, register *map[string][]interface{})  {
	VariableA,VariableB:=syntax.Tow(code)
	if register_structure.ExistVariable(VariableA, register){
		if register_structure.ExistVariable(VariableB, register){
			fmt.Println(register_structure.Cmp(VariableA,VariableB, register))
		}else{
			interpreter.Log("not found variable : "+VariableB,"add int function",2)
		}
	}else {
		interpreter.Log("not found variable : "+VariableA,"add int function",2)
	}
}