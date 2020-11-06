package list_structure

import (
	"../../interpreter"
	"fmt"
)

//get list value with index
//list(name,index)
func ListGet(ListName string,index int,ListMap *map[string][]interface{})interface{}{
	var list map[string][]interface{}=*ListMap
	//check exist list
	if ExistList(ListName,ListMap){
		//index is out of range
		if index>len(list[ListName])-1{
			return fmt.Sprintf("out of range len : %d  => %d",len(list[ListName])-1,index)
		}
		//item is nil
		if list[ListName][index]==nil{
			return "null"
		}
		//test search
		fmt.Printf("search {%v} index : {%d}\n", list[ListName][index], Search(ListName, list[ListName][index].(string), ListMap))
		return list[ListName][index]
	}else {
		interpreter.Log("not found list : "+ListName,"get list value",2)
	}
	return -1
}
