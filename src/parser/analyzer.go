package parser

import (
	"./syntax"
)

const STRING ="str"
const INT ="int"
func Core(code string,stack *map[string][]interface{})interface{}{

	if IsValid("print",code){
		syntax.Print(code,STRING)
		return true

	}else if IsValid("print_int",code){
		syntax.Print(code,INT)
		return true

	}else if IsValid("logic",code){
		syntax.Logic(code)
		return true

	}else if IsValid("math",code){
		syntax.Math(code)
		return true

	}else if IsValid("streql",code){
		syntax.StrEql(code)
		return true

	}else if IsValid("len",code){
		syntax.Len(code)
		return true

	}else if IsValid("find",code){
		syntax.Find(code)
		return true

	}else if IsValid("var_str",code) {
		syntax.Var(code,STRING,stack)
		return true

	}else if IsValid("var_int",code){
		syntax.Var(code,INT,stack)
		return true

	} else if IsValid("push_str",code){
		syntax.Push(code,STRING,stack)
		return true

	}else if IsValid("push_int",code){
		syntax.Push(code,INT,stack)
		return true

	} else if IsValid("pop",code){
		syntax.Pop(code,stack)
		return true

	}else if IsValid("rm",code){
		syntax.Rm(code,stack)
		return true

	}else if IsValid("add_int",code){
		syntax.Add(code,INT,stack)
		return true

	}else if IsValid("add_str",code){
		syntax.Add(code,STRING,stack)
		return true
	}else if IsValid("type",code){
		syntax.Stype(code,stack)
		return true
	}else if IsValid("id",code){
		syntax.Id(code,stack)
		return true
	}else if IsValid("ivar",code){
		return syntax.Ivar(code,stack)
	}else if IsValid("if_s1",code){
		syntax.IF(code)
		return true
	} else{
		return false
	}
}