package list_syntax

import (
	"../../../interpreter"
	"../../../interpreter/list-structure"
	"../../syntax"
	"fmt"
	"strconv"
)
//list push value syntax
func ListPush(code string,ListMap *map[string][]interface{}){
	name,index,value:=syntax.Three(code)
	fmt.Println(name,value,index)
	if index,err:=strconv.Atoi(index);err==nil{
		list_structure.ListPush(name,value,index,ListMap)
	}else{
		interpreter.Log("list push format error","list push",2)
	}
}
