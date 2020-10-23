package list_structure

import (
	"../../interpreter"
)

func ListPush(ListName,value string,index int,ListMap *map[string][]interface{}){
	if ExistList(ListName,ListMap){
		var list map[string][]interface{}=*ListMap
		list[ListName][index]=value
	}else {
		interpreter.Log("not found list","list push",2)
	}
}
