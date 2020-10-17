package syntax

import (
	"../../interpreter"
	"fmt"
)

//pop syntax
func Pop(code string,stack *map[string][]interface{})  {
	//check exist variable with name
	if interpreter.ExistVariable(code[4:len(code)-1],stack) {
		//pop([variable-name])
		fmt.Println(interpreter.GetValue(code[4:len(code)-1], stack))
	}else{
		interpreter.Log("not found variable","pop function",2)
	}
}
