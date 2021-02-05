package Variable_syntax

import (
	"../../../interpreter/register-structure"
	"../../syntax"
)
//create variable syntax
func Var(code,_type string, register *map[string][]interface{})  {
	VariableName:=syntax.One(code)
	if _type==STRING{
		register_structure.CreateVariable(VariableName,"str", register)
	}else if _type==INT{
		register_structure.CreateVariable(VariableName,"int", register)
	}
}