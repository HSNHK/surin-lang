package Variable_syntax

import (
	"../../../interpreter"
	"../../../interpreter/register-structure"
	"../../syntax"
)
//remove variable syntax
func Rm(code string, register *map[string][]interface{})  {
	VariableName:=syntax.One(code)
	if register_structure.DeleteVariable(VariableName, register)!=true{
		interpreter.Log("not found variable","rm function",2)
	}
}
