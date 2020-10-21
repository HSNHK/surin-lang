package parser

import (
	"./syntax"
	"./syntax/register-syntax"
	"./syntax/list-syntax"
)
//value type
const STRING ="str"
const INT ="int"
func Core(code string, register *map[string][]interface{}){
	//print("hello")
	if IsValid("print",code){
		syntax.Print(code,STRING)
	//print(1234)
	}else if IsValid("print_int",code){
		syntax.Print(code,INT)
	//logic(12,>,11)
	}else if IsValid("logic",code){
		syntax.Logic(code)
	//math(2,**,2)
	}else if IsValid("math",code){
		syntax.Math(code)
	//streql("hello","god")
	}else if IsValid("streql",code){
		syntax.StrEql(code)
	//len("hello..")
	}else if IsValid("len",code){
		syntax.Len(code)
	//len("hello","h")
	}else if IsValid("find",code){
		syntax.Find(code)
	//var("variable-name",str)
	}else if IsValid("var_str",code) {
		register_syntax.Var(code,STRING, register)
	//var("variable-name",int)
	}else if IsValid("var_int",code){
		register_syntax.Var(code,INT, register)
	//push("variable-name","hello")
	} else if IsValid("push_str",code){
		register_syntax.Push(code,STRING, register)
	//push("variable-name",1234)
	}else if IsValid("push_int",code){
		register_syntax.Push(code,INT, register)
	//pop(variable-name)
	} else if IsValid("pop",code){
		register_syntax.Pop(code, register)
	//rm(variable-name)
	}else if IsValid("rm",code){
		register_syntax.Rm(code, register)
	//add("variable-name",1234)
	}else if IsValid("add_int",code){
		register_syntax.Add(code,INT, register)
	//add("variable-name","hello")
	}else if IsValid("add_str",code){
		register_syntax.Add(code,STRING, register)
	//type(variable-name)
	}else if IsValid("type",code){
		register_syntax.Stype(code, register)
	//id(variable-name)
	}else if IsValid("id",code){
		register_syntax.Id(code, register)
	//ivar(variable-name)

	/*
	}else if IsValid("ivar",code){
		return syntax.Ivar(code,register)
	 */

	//if(1,>,2)?print("yes"):print("no")
	}else if IsValid("if_s1",code){
		syntax.IF(code)
	}else if IsValid("cmp",code){
		register_syntax.Cmp(code, register)
	}else if IsValid("list",code){

	}
}