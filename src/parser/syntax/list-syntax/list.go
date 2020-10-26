package list_syntax

import (
	"../../../interpreter"
	"../../../interpreter/list-structure"
	"fmt"
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
//show all item
//list(name)
func ShowList(code string,ListMap *map[string][]interface{})  {
	//list([name])
	fmt.Println(list_structure.List(code[5:len(code)-1],ListMap))
}