package list_structure

import "fmt"
//bauble sort function
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
//binary search function
func Search(name,Value string,ListMap *map[string][]interface{})int{
	var list map[string][]interface{}=*ListMap
	//item index found
	found:=0
	//start list index
	start:=0
	//end list index
	end:=0
	//get list size
	for i:=0;i<len(list[name]);i++{
		//filter nil item
		if list[name][i]!=nil{
			end++
		}
	}
	fmt.Println("len : ",end)
	for  start<=end{
		//list size / 2
		mid:=(start+end)/2
		//compare value with list[len/2]
		if fmt.Sprintf("%v",list[name][mid])==Value{
			found=mid
			break
		//check list[*:]
		}else if fmt.Sprintf("%v",list[name][mid])>Value {
			end=start-1
		//check list[:*]
		}else if fmt.Sprintf("%v",list[name][mid])<Value{
			start=mid+1
		}
	}
	//not fund = -1
	return found
}
