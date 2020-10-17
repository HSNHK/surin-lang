package syntax

import (
	"../../interpreter"
	"strings"
)

//push function
func Push(code,_type string,stack *map[string][]interface{})  {
	if _type==STRING{
		//push("variable-name"[,]"hello")
		codeSplit:=strings.Split(code,",")
		//check exist variable with name
		if interpreter.ExistVariable(codeSplit[0][6:len(codeSplit[0])-1],stack){
			//push(["variable-name"],["hello"])
			interpreter.PushStack(codeSplit[0][6:len(codeSplit[0])-1],STRING,
				codeSplit[1][1:len(codeSplit[1])-2],stack)
		}else{
			interpreter.Log("not found variable","push str function",2)
		}
	}else if _type ==INT{
		//push("variable-name"[,]1234)
		codeSplit:=strings.Split(code,",")
		//check exist variable with name
		if interpreter.ExistVariable(codeSplit[0][6:len(codeSplit[0])-1],stack){
			//push(["variable-name"],[01234])
			interpreter.PushStack(codeSplit[0][6:len(codeSplit[0])-1],INT,
				codeSplit[1][:len(codeSplit[1])-1],stack)
		}else{
			interpreter.Log("not found variable","push int function",2)
		}
	}
}