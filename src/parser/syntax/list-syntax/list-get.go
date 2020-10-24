package list_syntax

import (
	"../../../interpreter"
	"../../../interpreter/list-structure"
	"fmt"
	"strconv"
	"strings"
)

//list get value syntax
//list(name,index)
func ListGet(code string,ListMap *map[string][]interface{}){
	//list(name[,]index)
	codesplit:=strings.Split(code,",")
	//list(name,[index])
	//convert string to int
	index,err:=strconv.Atoi(codesplit[1][:len(codesplit[1])-1])
	if err==nil{
		//list([name],[index])
		fmt.Println(list_structure.ListGet(codesplit[0][5:],index,ListMap))
	}else{
		interpreter.Log("index error","list get value syntax",2)
	}
}
