package list_syntax

import (
	"../../../interpreter"
	"../../../interpreter/list-structure"
	"../../syntax"
	"fmt"
	"strconv"
)
//creat list syntax
func List(code string,ListMap *map[string][]interface{})  {
	name,size:=syntax.Initialise(code)
	if s,err:=strconv.Atoi(size);err==nil{
		list_structure.ListCreate(name,s,ListMap)
	}else {
		interpreter.Log("list size error","create list",3)
	}
}
//show all item
func ShowList(code string,ListMap *map[string][]interface{})  {
	//list([name])
	fmt.Println(list_structure.List(syntax.One(code),ListMap))
}
//search method syntax
func Search(code string,ListMap *map[string][]interface{})  {
	name,value:=syntax.Tow(code)
	fmt.Println(list_structure.Search(name,value,ListMap))
}
//sort list
func Sort(code string,ListMap *map[string][]interface{})  {
	list_structure.Sort(syntax.One(code),ListMap)
}
//delete list
func Delete(code string,ListMap *map[string][]interface{}){
	list_structure.Delete(syntax.One(code),ListMap)
}