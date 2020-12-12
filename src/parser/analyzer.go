package parser

import (
	"./syntax"
	"./syntax/register-syntax"
	"fmt"
)
//value type
const STRING ="str"
const INT ="int"
func Core(code string, register,list *map[string][]interface{}){

	if IsValidNameSpace("List",code){
		ListAnalyzer(code,list)
	//print("hello")
	}else if IsValid("Global","print",code){
		syntax.Print(code,STRING)

	//print(1234)
	}else if IsValid("Global","print_int",code){
		syntax.Print(code,INT)

	//logic(12,>,11)
	}else if IsValid("Global","logic",code){
		syntax.Logic(code)

	//math(2,**,2)
	}else if IsValid("Global","math",code){
		syntax.Math(code)

	//streql("hello","god")
	}else if IsValid("Global","streql",code){
		syntax.StrEql(code)

	//len("hello..")
	}else if IsValid("Global","len",code){
		syntax.Len(code)

	//len("hello","h")
	}else if IsValid("Global","find",code){
		syntax.Find(code)

	//var("variable-name",str)
	}else if IsValid("Global","var_str",code) {
		register_syntax.Var(code,STRING, register)

	//var("variable-name",int)
	}else if IsValid("Global","var_int",code){
		register_syntax.Var(code,INT, register)

	//push("variable-name","hello")
	} else if IsValid("Global","push_str",code){
		register_syntax.Push(code,STRING, register)

	//push("variable-name",1234)
	}else if IsValid("Global","push_int",code){
		register_syntax.Push(code,INT, register)

	//pop(variable-name)
	} else if IsValid("Global","pop",code){
		register_syntax.Pop(code, register)

	//rm(variable-name)
	}else if IsValid("Global","rm",code){
		register_syntax.Rm(code, register)

	//add("variable-name",1234)
	}else if IsValid("Global","add_int",code){
		register_syntax.Add(code,INT, register)

	//add("variable-name","hello")
	}else if IsValid("Global","add_str",code){
		register_syntax.Add(code,STRING, register)

	//type(variable-name)
	}else if IsValid("Global","type",code){
		register_syntax.Stype(code, register)

	//id(variable-name)
	}else if IsValid("Global","id",code){
		register_syntax.Id(code, register)

	//ivar(variable-name)

	/*
	}else if IsValid("ivar",code){
		return syntax.Ivar(code,register)
	 */

	//if(1,>,2)?print("yes"):print("no")
	}else if IsValid("Global","if_s1",code){
		syntax.IF(code)

	//cmp(variable-name-1,variable-name-2)
	}else if IsValid("Global","cmp",code){
		register_syntax.Cmp(code, register)
	}else{
		fmt.Println("command fail ..!")
	}
}