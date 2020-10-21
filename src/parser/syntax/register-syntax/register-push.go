package syntax

import (
	"../../interpreter"
	"../../interpreter/register-structure"
	"strings"
)

//push function
func Push(code,_type string, register *map[string][]interface{})  {
	if _type==STRING{
		//push("variable-name"[,]"hello")
		codeSplit:=strings.Split(code,",")
		//check exist variable with name
		if register_structure.ExistVariable(codeSplit[0][6:len(codeSplit[0])-1], register){
			//push(["variable-name"],["hello"])
			register_structure.PushStack(codeSplit[0][6:len(codeSplit[0])-1],STRING,
				codeSplit[1][1:len(codeSplit[1])-2], register)
		}else{
			interpreter.Log("not found variable","push str function",2)
		}
	}else if _type ==INT{
		//push("variable-name"[,]1234)
		codeSplit:=strings.Split(code,",")
		//check exist variable with name
		if register_structure.ExistVariable(codeSplit[0][6:len(codeSplit[0])-1], register){
			//push(["variable-name"],[01234])
			register_structure.PushStack(codeSplit[0][6:len(codeSplit[0])-1],INT,
				codeSplit[1][:len(codeSplit[1])-1], register)
		}else{
			interpreter.Log("not found variable","push int function",2)
		}
	}
}