package register_syntax

import (
	"../../../interpreter"
	"../../../interpreter/register-structure"
	"strings"
)

//add to variable syntax
func Add(code,_type string, register *map[string][]interface{})  {
	if _type==STRING{
		//add("variable-name"[,]"hello")
		codeSplit:=strings.Split(code,",")
		//check exist variable
		if register_structure.ExistVariable(codeSplit[0][5:len(codeSplit[0])-1], register) {
			//add(["variable-name"],["hello"])
			register_structure.AddToVariable(codeSplit[0][5:len(codeSplit[0])-1], "str", codeSplit[1][1:len(codeSplit[1])-2], register)
		}else{
			interpreter.Log("not found variable","add str function",2)
		}
	}else if _type==INT{
		//add("variable-name"[,]1234)
		codeSplit:=strings.Split(code,",")
		//check exist variable
		if register_structure.ExistVariable(codeSplit[0][5:len(codeSplit[0])-1], register) {
			//add(["variable-name"],[1234])
			register_structure.AddToVariable(codeSplit[0][5:len(codeSplit[0])-1], "int", codeSplit[1][:len(codeSplit[1])-1], register)
		}else{
			interpreter.Log("not found variable","add int function",2)
		}
	}
}
