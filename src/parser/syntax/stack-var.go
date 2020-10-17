package syntax

import "../../interpreter"

//create variable syntax
func  Var(code,_type string,stack *map[string][]interface{})  {
	if _type==STRING{
		//var(["variable-name"],[str])
		interpreter.CreateVariable(code[4:len(code)-5],"str",stack)
	}else if _type==INT{
		//var(["variable-name"],[int])
		interpreter.CreateVariable(code[4:len(code)-5],"int",stack)
	}
}