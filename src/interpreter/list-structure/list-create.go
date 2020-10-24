package list_structure

import (
	"../../interpreter"
)
//create list
//list(name:size)
func ListCreate(listName string,size int,ListMap *map[string][]interface{}){
	//check exist list
	if ExistList(listName,ListMap){
		interpreter.Log("found list : "+listName,"list create",2)
	}else {
		var list map[string][]interface{}=*ListMap
		list[listName]=make([]interface{},size)
	}
}