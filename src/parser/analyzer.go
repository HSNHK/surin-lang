package parser

import (
	"./syntax"
	"./syntax/list-syntax"
	"./syntax/register-syntax"
	"fmt"
)
//value type
const STRING ="str"
const INT ="int"
func Core(code string, register,list *map[string][]interface{}){
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
	//cmp(variable-name-1,variable-name-2)
	}else if IsValid("cmp",code){
		register_syntax.Cmp(code, register)
	//list(name:size)
	}else if IsValid("list",code){
		list_syntax.List(code,list)
	//list(name,index,value)
	}else if IsValid("list-push",code){
		list_syntax.ListPush(code,list)
	//list(name,index)
	}else if IsValid("list-get",code) {
		list_syntax.ListGet(code,list)
	//list(list-name)
	}else if IsValid("show-all-list-item",code) {
		list_syntax.ShowList(code,list)
	}else if IsValid("list-search",code) {
		list_syntax.Search(code,list)
	}else{
		fmt.Println("command fail ..!")
	}
}