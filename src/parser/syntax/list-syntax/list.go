package list_syntax

import (
	"../../../interpreter"
	"../../../interpreter/list-structure"
	"strconv"
	"strings"
)

func List(code string,ListMap *map[string][]interface{})  {
	codesplit:=strings.Split(code,":")
	size,err:=strconv.Atoi(codesplit[1][:len(codesplit[1])-1])
	if err==nil{
		list_structure.ListCreate(codesplit[0][5:],size,ListMap)
	}else {
		interpreter.Log("list size error","create list",3)
	}
}

