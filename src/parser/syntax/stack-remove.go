package syntax

import "../../interpreter"

//remove variable syntax
func Rm(code string,stack *map[string][]interface{})  {
	//rm([variable-name])
	if interpreter.DeleteVariable(code[3:len(code)-1],stack)!=true{
		interpreter.Log("not found variable","rm function",2)
	}
}
