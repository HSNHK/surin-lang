package list_syntax

import (
	"../../../interpreter"
	"../../../interpreter/list-structure"
	"strconv"
	"strings"
)

//list push value syntax
//list(name,index,value)
func ListPush(code string,ListMap *map[string][]interface{}){
	//list(name[,]index[,]value)
	codesplit:=strings.Split(code,",")
	//list(name,[index],value)
	//convert string to int
	index,err:=strconv.Atoi(codesplit[1])
	if err==nil{
		//list(name,index,[value])
		value:=codesplit[2][:len(codesplit[2])-1]
		//list([name],[index],[value])
		list_structure.ListPush(codesplit[0][5:],value,index,ListMap)
	}else{
		interpreter.Log("list push format error","list push",2)
	}
}
