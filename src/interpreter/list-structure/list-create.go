package list_structure

import (
	"../../interpreter"
)

func ListCreate(listName string,ListMap *map[string][]interface{}){
	if ExistList(listName,ListMap){
		interpreter.Log("found list : "+listName,"list create",2)
	}else {
		var list map[string][]interface{}=*ListMap
		list[listName]=make([]interface{},100)
	}
}