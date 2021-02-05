package Variable_syntax

import (
	"../../../interpreter"
	"../../../interpreter/register-structure"
	"../../syntax"
)

//push function
func Push(code,_type string, register *map[string][]interface{})  {
	if _type==STRING{
		VariableName,value:=syntax.Tow(code)
		//check exist variable with name
		if register_structure.ExistVariable(VariableName, register){
			register_structure.PushStack(VariableName,STRING,value, register)
		}else{
			interpreter.Log("not found variable","push str function",2)
		}
	}else if _type ==INT{
		VariableName,value:=syntax.Tow(code)
		//check exist variable with name
		if register_structure.ExistVariable(VariableName, register){
			register_structure.PushStack(VariableName,INT,value, register)
		}else{
			interpreter.Log("not found variable","push int function",2)
		}
	}
}