package list_syntax

import (
	"../../../interpreter"
	"../../../interpreter/list-structure"
	"fmt"
	"strconv"
	"strings"
)

func ListPush(code string,ListMap *map[string][]interface{}){
	fmt.Println(code)
	codesplit:=strings.Split(code,",")
	name:=codesplit[0][5:]
	index,err:=strconv.Atoi(codesplit[1])
	if err==nil{
		value:=codesplit[2][:len(codesplit[2])-1]
		list_structure.ListPush(name,value,index,ListMap)
	}else{
		interpreter.Log("list push format error","list push",2)
	}
}
