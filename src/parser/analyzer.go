package parser

import (
	"./syntax"
)
//value type
const STRING ="str"
const INT ="int"
func Core(code string,stack *map[string][]interface{})interface{}{
	//print("hello")
	if IsValid("print",code){
		syntax.Print(code,STRING)
		return true
	//print(1234)
	}else if IsValid("print_int",code){
		syntax.Print(code,INT)
		return true
	//logic(12,>,11)
	}else if IsValid("logic",code){
		syntax.Logic(code)
		return true
	//math(2,**,2)
	}else if IsValid("math",code){
		syntax.Math(code)
		return true
	//streql("hello","god")
	}else if IsValid("streql",code){
		syntax.StrEql(code)
		return true
	//len("hello..")
	}else if IsValid("len",code){
		syntax.Len(code)
		return true
	//len("hello","h")
	}else if IsValid("find",code){
		syntax.Find(code)
		return true
	//var("variable-name",str)
	}else if IsValid("var_str",code) {
		syntax.Var(code,STRING,stack)
		return true
	//var("variable-name",int)
	}else if IsValid("var_int",code){
		syntax.Var(code,INT,stack)
		return true
	//push("variable-name","hello")
	} else if IsValid("push_str",code){
		syntax.Push(code,STRING,stack)
		return true
	//push("variable-name",1234)
	}else if IsValid("push_int",code){
		syntax.Push(code,INT,stack)
		return true
	//pop(variable-name)
	} else if IsValid("pop",code){
		syntax.Pop(code,stack)
		return true
	//rm(variable-name)
	}else if IsValid("rm",code){
		syntax.Rm(code,stack)
		return true
	//add("variable-name",1234)
	}else if IsValid("add_int",code){
		syntax.Add(code,INT,stack)
		return true
	//add("variable-name","hello")
	}else if IsValid("add_str",code){
		syntax.Add(code,STRING,stack)
		return true
	//type(variable-name)
	}else if IsValid("type",code){
		syntax.Stype(code,stack)
		return true
	//id(variable-name)
	}else if IsValid("id",code){
		syntax.Id(code,stack)
		return true
	//ivar(variable-name)
	}else if IsValid("ivar",code){
		return syntax.Ivar(code,stack)
	//if(1,>,2)?print("yes"):print("no")
	}else if IsValid("if_s1",code){
		syntax.IF(code)
		return true
	} else{
		return false
	}
}