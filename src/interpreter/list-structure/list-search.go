package list_structure

import "fmt"

func sort(name string,ListMap *map[string][]interface{}){
	var list=*ListMap
	for index:=len(list[name])-2;index>=0;index--{
		if fmt.Sprintf("%T",list[name][index])=="string" {
			for i := 0; i<=index; i++ {
				if list[name][i]!=nil {
					if fmt.Sprintf("%v", list[name][i]) >
					   fmt.Sprintf("%v", list[name][i+1]){
						//fmt.Println(list[name][i]," > ",list[name][i+1])
						next:=list[name][i+1]
						list[name][i+1]=list[name][i]
						list[name][i]=next
					}
				}
			}
		}else if fmt.Sprintf("%T",list[name][index])=="int"{
			for i := 0; i<=index; i++ {
				if list[name][i]!=nil {
					if fmt.Sprintf("%v", list[name][i]) >
						fmt.Sprintf("%v", list[name][i+1]){
						fmt.Println(list[name][i]," > ",list[name][i+1])
						temp:=list[name][i]
						list[name][i]=list[name][i+1]
						list[name][i+1]=temp
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
