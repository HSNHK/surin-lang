package Variable_syntax

import (
	"../../../interpreter"
	"../../../interpreter/register-structure"
	"../../syntax"
)

//add to variable syntax
func Add(code,_type string, register *map[string][]interface{})  {
	if _type==STRING{
		VariableName,Value:=syntax.Tow(code)
		//check exist variable
		if register_structure.ExistVariable(VariableName, register) {
			register_structure.AddToVariable(VariableName, "str", Value, register)
		}else{
			interpreter.Log("not found variable","add str function",2)
		}
	}else if _type==INT{
		VariableName,Value:=syntax.Tow(code)
		//check exist variable
		if register_structure.ExistVariable(VariableName, register) {
			register_structure.AddToVariable(VariableName, "int",Value, register)
		}else{
			interpreter.Log("not found variable","add int function",2)
		}
	}
}
