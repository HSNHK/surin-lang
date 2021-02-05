package Variable_syntax

import (
	"../../../interpreter"
	"../../../interpreter/register-structure"
	"../../syntax"
	"fmt"
)

//pop syntax
func Pop(code string, register *map[string][]interface{})  {
	//check exist variable with name
	VariableName :=syntax.One(code)
	if register_structure.ExistVariable(VariableName, register) {
		fmt.Println(register_structure.GetValue(VariableName, register))
	}else{
		interpreter.Log("not found variable","pop function",2)
	}
}
