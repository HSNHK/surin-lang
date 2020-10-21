package register_syntax

import (
	"../../../interpreter"
	"../../../interpreter/register-structure"
)

//remove variable syntax
func Rm(code string, register *map[string][]interface{})  {
	//rm([variable-name])
	if register_structure.DeleteVariable(code[3:len(code)-1], register)!=true{
		interpreter.Log("not found variable","rm function",2)
	}
}
