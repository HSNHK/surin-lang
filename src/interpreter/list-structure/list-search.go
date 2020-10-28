package list_structure

import "fmt"

func sort(name string,ListMap *map[string][]interface{}){
	var list=*ListMap
	for index,value:=range list[name]{
		if fmt.Sprintf("%T",value)=="string"{
			if fmt.Sprintf("%T",list[name][index+1])=="string"{
				if fmt.Sprintf("%v",value)[0]>
					fmt.Sprintf("%v",list[name][index+1])[0]{
					next:=list[name][index+1]
					list[name][index+1]=value
					list[name][index]=next
				}
			}else if fmt.Sprintf("%T",list[name][index+1])=="int"{
				next:=list[name][index+1]
				list[name][index+1]=value
				list[name][index]=next
			}
		}else if fmt.Sprintf("%T",value) =="int"{
			for i:=index;i<=cap(list[name]);i++{
				if fmt.Sprintf("%T",list[name][i+1])=="int"{
					if fmt.Sprintf("%v",value)>
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

func Search(value string,ListMap *map[string][]interface{}){

}
