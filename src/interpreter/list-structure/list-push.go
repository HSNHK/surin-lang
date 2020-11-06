package list_structure

import (
	"../../interpreter"
	"fmt"
)

//push value to list
//list(name,index,value)
func ListPush(ListName,value string,index int,ListMap *map[string][]interface{}){
	if ExistList(ListName,ListMap){
		var list map[string][]interface{}=*ListMap
		//index is out of range
		if index>len(list[ListName])-1{
			fmt.Println(fmt.Sprintf("out of range len : %d  => %d",len(list[ListName])-1,index))
		}else {
			list[ListName][index] = value
		}
	}else {
		interpreter.Log("not found list","list push",2)
	}
}