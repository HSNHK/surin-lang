package list_structure

import "../../interpreter"

//get list value with index
//list(name,index)
func ListGet(ListName string,index int,ListMap *map[string][]interface{})interface{}{
	var list map[string][]interface{}=*ListMap
	//check exist list
	if ExistList(ListName,ListMap){
		return list[ListName][index]
	}else {
		interpreter.Log("not found list : "+ListName,"get list value",2)
	}
	return -1
}
