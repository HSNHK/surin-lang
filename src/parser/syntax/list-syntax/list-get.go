package list_syntax

import (
	"../../../interpreter"
	"../../../interpreter/list-structure"
	"fmt"
	"strconv"
	"strings"
)

func ListGet(code string,ListMap *map[string][]interface{}){
	codesplit:=strings.Split(code,",")
	index,err:=strconv.Atoi(codesplit[1][:len(codesplit[1])-1])
	if err==nil{
		fmt.Println(list_structure.ListGet(codesplit[0][5:],index,ListMap))
	}else{
		interpreter.Log("index error","list get value syntax",2)
	}
}
