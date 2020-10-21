package syntax

import (
	"../../interpreter/register-structure"
)

//create variable syntax
func  Var(code,_type string, register *map[string][]interface{})  {
	if _type==STRING{
		//var(["variable-name"],[str])
		register_structure.CreateVariable(code[4:len(code)-5],"str", register)
	}else if _type==INT{
		//var(["variable-name"],[int])
		register_structure.CreateVariable(code[4:len(code)-5],"int", register)
	}
}