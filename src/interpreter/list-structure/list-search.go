package list_structure

import "fmt"

func sort(name string,ListMap *map[string][]interface{}){
	var list=*ListMap
	for index,value:=range list[name]{
		if fmt.Sprintf("%T",value)=="string" {
			for i := index; i < len(list[name]); i++ {
				if list[name][i]!=nil {
					fmt.Println(list[name][i])
					if fmt.Sprintf("%v", list[name][i]) >
					   fmt.Sprintf("%v", list[name][i+1]){
						next := list[name][i+1]
						list[name][i+1] = list[name][i]
						list[name][i] = next
					}
				}
			}
		}else if fmt.Sprintf("%T",value) =="int"{
			for i:=index;i<=cap(list[name]);i++{
				if fmt.Sprintf("%T",list[name][i+1])=="int"{
					if fmt.Sprintf("%v",list[name][i])>
						fmt.Sprintf("%v",list[name][i+1]){
						next:=list[name][i+1]
						list[name][i+1]=value
						list[name][i]=next
					}
				}
			}
		}
	}
}

func Search(name,Value string,ListMap *map[string][]interface{})int{
	var list=*ListMap
	for index,value:=range list[name]{
		if value==Value{
			return index
		}
	}
	return -1
}
