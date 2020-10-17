package syntax

import (
	"../../interpreter"
	"strings"
)

//add to variable syntax
func Add(code,_type string,stack *map[string][]interface{})  {
	if _type==STRING{
		//add("variable-name"[,]"hello")
		codeSplit:=strings.Split(code,",")
		//check exist variable
		if interpreter.ExistVariable(codeSplit[0][5:len(codeSplit[0])-1],stack) {
			//add(["variable-name"],["hello"])
			interpreter.AddToVariable(codeSplit[0][5:len(codeSplit[0])-1], "str", codeSplit[1][1:len(codeSplit[1])-2], stack)
		}else{
			interpreter.Log("not found variable","add str function",2)
		}
	}else if _type==INT{
		//add("variable-name"[,]1234)
		codeSplit:=strings.Split(code,",")
		//check exist variable
		if interpreter.ExistVariable(codeSplit[0][5:len(codeSplit[0])-1],stack) {
			//add(["variable-name"],[1234])
			interpreter.AddToVariable(codeSplit[0][5:len(codeSplit[0])-1], "int", codeSplit[1][:len(codeSplit[1])-1], stack)
		}else{
			interpreter.Log("not found variable","add int function",2)
		}
	}
}
