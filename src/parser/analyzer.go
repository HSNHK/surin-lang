package parser

import (
	"./syntax"
)
//value type
const STRING ="str"
const INT ="int"
func Core(code string,stack *map[string][]interface{}){
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
		syntax.Var(code,STRING,stack)
	//var("variable-name",int)
	}else if IsValid("var_int",code){
		syntax.Var(code,INT,stack)
	//push("variable-name","hello")
	} else if IsValid("push_str",code){
		syntax.Push(code,STRING,stack)
	//push("variable-name",1234)
	}else if IsValid("push_int",code){
		syntax.Push(code,INT,stack)
	//pop(variable-name)
	} else if IsValid("pop",code){
		syntax.Pop(code,stack)
	//rm(variable-name)
	}else if IsValid("rm",code){
		syntax.Rm(code,stack)
	//add("variable-name",1234)
	}else if IsValid("add_int",code){
		syntax.Add(code,INT,stack)
	//add("variable-name","hello")
	}else if IsValid("add_str",code){
		syntax.Add(code,STRING,stack)
	//type(variable-name)
	}else if IsValid("type",code){
		syntax.Stype(code,stack)
	//id(variable-name)
	}else if IsValid("id",code){
		syntax.Id(code,stack)
	//ivar(variable-name)

	/*
	}else if IsValid("ivar",code){
		return syntax.Ivar(code,stack)
	 */

	//if(1,>,2)?print("yes"):print("no")
	}else if IsValid("if_s1",code){
		syntax.IF(code)
	}else if IsValid("cmp",code){
		syntax.Cmp(code,stack)
	}
}