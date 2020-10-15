package syntax

import (
	"../../interpreter"
	"fmt"
	"strings"
)
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
//remove variable syntax
func Rm(code string,stack *map[string][]interface{})  {
	//rm([variable-name])
	if interpreter.DeleteVariable(code[3:len(code)-1],stack)!=true{
		interpreter.Log("not found variable","rm function",2)
	}
}
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
//get variable type syntax
func Stype(code string,stack *map[string][]interface{})  {
	//check exist variable
	if interpreter.ExistVariable(code[5:len(code)-1],stack){
		//type([variable-name])
		fmt.Println(interpreter.Type(code[5:len(code)-1],stack))
	}else{
		interpreter.Log("not found variable","add int function",2)
	}
}
//get variable id syntax
func Id(code string,stack *map[string][]interface{})  {
	//check exist variable
	if interpreter.ExistVariable(code[3:len(code)-1],stack){
		//id([variable-name])
		fmt.Println(interpreter.Id(code[3:len(code)-1],stack))
	}else{
		interpreter.Log("not found variable","add int function",2)
	}
}