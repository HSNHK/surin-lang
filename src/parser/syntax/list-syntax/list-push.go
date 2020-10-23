package list_syntax

import (
	"../../../interpreter"
	"../../../interpreter/list-structure"
	"strconv"
	"strings"
)

func ListPush(code string,ListMap *map[string][]interface{}){
	codesplit:=strings.Split(code,",")
	index,err:=strconv.Atoi(codesplit[1])
	if err==nil{
		value:=codesplit[2][:len(codesplit[2])-1]
		list_structure.ListPush(codesplit[0][5:],value,index,ListMap)
	}else{
		interpreter.Log("list push format error","list push",2)
	}
}
