package register_syntax

import (
	"../../../interpreter"
	"../../../interpreter/register-structure"
	"fmt"
)

//pop syntax
func Pop(code string, register *map[string][]interface{})  {
	//check exist variable with name
	if register_structure.ExistVariable(code[4:len(code)-1], register) {
		//pop([variable-name])
		fmt.Println(register_structure.GetValue(code[4:len(code)-1], register))
	}else{
		interpreter.Log("not found variable","pop function",2)
	}
}
