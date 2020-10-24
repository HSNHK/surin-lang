package list_syntax

import (
	"../../../interpreter"
	"../../../interpreter/list-structure"
	"strconv"
	"strings"
)

//creat list syntax
//list(name:size)
func List(code string,ListMap *map[string][]interface{})  {
	//list(name[:]size)
	codesplit:=strings.Split(code,":")
	//list(name:[size])
	//convert string to int
	size,err:=strconv.Atoi(codesplit[1][:len(codesplit[1])-1])
	if err==nil{
		//list([name]:[size])
		list_structure.ListCreate(codesplit[0][5:],size,ListMap)
	}else {
		interpreter.Log("list size error","create list",3)
	}
}

