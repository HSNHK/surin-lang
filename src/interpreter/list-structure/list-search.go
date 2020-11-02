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
	var list map[string][]interface{}=*ListMap
	found:=0
	start:=0
	end:=len(list[name])-1
	for  start<=end{
		mid:=(start+end)/2
		if fmt.Sprintf("%v",list[name][mid])==Value{
			found=mid
			break
		}else if fmt.Sprintf("%v",list[name][mid])>Value {
			end=start-1
		}else if fmt.Sprintf("%v",list[name][mid])<Value{
			start=mid+1
		}
	}
	return found
}
