package list_structure

import (
	"../../interpreter"
	"fmt"
)

func ListPush(ListName,value string,index int,ListMap *map[string][]interface{}){
	if ExistList(ListName,ListMap){
		var list map[string][]interface{}=*ListMap
		list[ListName][index]=value
		fmt.Println("ok push")
	}else {
		interpreter.Log("not found list","list push",2)
	}
}
