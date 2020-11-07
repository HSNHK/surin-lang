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
//search method syntax
//list.search(list-name,value)
func Search(code string,ListMap *map[string][]interface{})  {
	//list.search[(list-name,value)]
	codesplit_step_1:=strings.Split(code,"(")
	//list.search([list-name,value)]
	codesplit_step_2:=strings.Split(codesplit_step_1[1],",")
	//remove list.search(list-name,value[)]
	codesplit_step_2[1]=codesplit_step_2[1][:len(codesplit_step_2[1])-1]
	fmt.Println(list_structure.Search(codesplit_step_2[0],codesplit_step_2[1],ListMap))
}