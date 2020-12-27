package parser

import (
	"./syntax"
	"./syntax/Variable-syntax"
	"fmt"
)
//value type
const STRING ="str"
const INT ="int"
func MainParser(code string, VariablesStorageSpace, ListsStorageSpace *map[string][]interface{}){
	//Check the code namespace to navigate correctly to your class
	if IsValidNameSpace("List",code){
		ListAnalyzer(code, ListsStorageSpace)

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
		Variable_syntax.Var(code,STRING, VariablesStorageSpace)

	//var("variable-name",int)
	}else if IsValid("Global","var_int",code){
		Variable_syntax.Var(code,INT, VariablesStorageSpace)

	//push("variable-name","hello")
	} else if IsValid("Global","push_str",code){
		Variable_syntax.Push(code,STRING, VariablesStorageSpace)

	//push("variable-name",1234)
	}else if IsValid("Global","push_int",code){
		Variable_syntax.Push(code,INT, VariablesStorageSpace)

	//pop(variable-name)
	} else if IsValid("Global","pop",code){
		Variable_syntax.Pop(code, VariablesStorageSpace)

	//rm(variable-name)
	}else if IsValid("Global","rm",code){
		Variable_syntax.Rm(code, VariablesStorageSpace)

	//add("variable-name",1234)
	}else if IsValid("Global","add_int",code){
		Variable_syntax.Add(code,INT, VariablesStorageSpace)

	//add("variable-name","hello")
	}else if IsValid("Global","add_str",code){
		Variable_syntax.Add(code,STRING, VariablesStorageSpace)

	//type(variable-name)
	}else if IsValid("Global","type",code){
		Variable_syntax.Stype(code, VariablesStorageSpace)

	//id(variable-name)
	}else if IsValid("Global","id",code){
		Variable_syntax.Id(code, VariablesStorageSpace)

	//if(1,>,2)?print("yes"):print("no")
	}else if IsValid("Global","if_s1",code){
		syntax.IF(code)

	//cmp(variable-name-1,variable-name-2)
	}else if IsValid("Global","cmp",code){
		Variable_syntax.Cmp(code, VariablesStorageSpace)
	}else{
		fmt.Println("command fail ..!")
	}
}