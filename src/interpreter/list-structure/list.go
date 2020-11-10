package list_structure

import (
	"../../interpreter"
	"fmt"
)

//check exist list with name
func ExistList(listName string,ListMap *map[string][]interface{}) bool{
	//var list map[string][]interface{}=*ListMap
	for key:=range *ListMap{
		if listName==key{
			return true
		}
	}
	return false
}
//show all item
//list(name)
func List(name string,ListMap *map[string][]interface{}) string {
	var list map[string][]interface{}=*ListMap
	//Sort(name,&list)
	//check exist list
	if ExistList(name,ListMap){
		for key:=range list{
			//check key name
			if key==name{
				//	[
				//  	 index : value
				//	]
				data:="[\n"
				for index,value:=range list[name]{
					//check != null
					if value!=nil{
						data+=fmt.Sprintf("	%d : %v\n",index,value)
					}
				}
				data+="] by list : "+name
				return data
			}
		}
		//not found list or empty list
		return "-1"
	}else {
		//not found list
		interpreter.Log("not fund list : "+name,"Show list",2)
		return "-1"
	}
}
//delete list
//list.del(list-name)
func Delete(name string,ListMap *map[string][]interface{}){
	if ExistList(name,ListMap){
		delete(*ListMap,name)
	}else {
		interpreter.Log("not found list : "+name,"delete list",2)
	}
}