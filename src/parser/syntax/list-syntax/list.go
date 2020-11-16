package list_syntax

import (
	"../../../interpreter"
	"../../../interpreter/list-structure"
	"../../syntax"
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
	codesplitStep1 :=strings.Split(code,"(")
	//list.search([list-name,value)]
	codesplitStep2 :=strings.Split(codesplitStep1[1],",")
	//remove list.search(list-name,value[)]
	codesplitStep2[1]= codesplitStep2[1][:len(codesplitStep2[1])-1]
	fmt.Println(codesplitStep2[0])
	fmt.Println(codesplitStep2[1])
	fmt.Println(list_structure.Search(codesplitStep2[0], codesplitStep2[1],ListMap))
}
//list.sort(list-name)
func Sort(code string,ListMap *map[string][]interface{})  {
	list_structure.Sort(syntax.One(code),ListMap)
}
//delete list
//list.delete(list-name)
func Delete(code string,ListMap *map[string][]interface{}){
	list_structure.Delete(syntax.One(code),ListMap)
}