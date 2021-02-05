package list_syntax

import (
	"../../../interpreter"
	"../../../interpreter/list-structure"
	"../../syntax"
	"fmt"
	"strconv"
)
//list get value syntax
func ListGet(code string,ListMap *map[string][]interface{}){
	name,index:=syntax.Tow(code)
	if index,err:=strconv.Atoi(index);err==nil{
		fmt.Println(list_structure.ListGet(name,index,ListMap))
	}else{
		interpreter.Log("index error","list get value syntax",2)
	}
}
