package list_structure

import "fmt"


//bubble sort function
//https://en.wikipedia.org/wiki/Bubble_sort
func sort(name string,ListMap *map[string][]interface{}){
	var list=*ListMap
	for index:=len(list[name])-2;index>=0;index--{
		//check item type
		if fmt.Sprintf("%T",list[name][index])=="string" {
			//compare item with all item
			for i := 0; i<=index; i++ {
				//check item != null
				if list[name][i]!=nil {
					//item1 > item2
					if fmt.Sprintf("%v", list[name][i]) >
					   fmt.Sprintf("%v", list[name][i+1]){
						//fmt.Println(list[name][i]," > ",list[name][i+1])
						//next = item2
						next:=list[name][i+1]
						//item2 = item1
						list[name][i+1]=list[name][i]
						//item1 = next
						list[name][i]=next
					}
				}
			}
		}else if fmt.Sprintf("%T",list[name][index])=="int"{
			for i := 0; i<=index; i++ {
				//check item != null
				if list[name][i]!=nil {
					//item1 > item2
					if fmt.Sprintf("%v", list[name][i]) >
					   fmt.Sprintf("%v", list[name][i+1]){
						//fmt.Println(list[name][i]," > ",list[name][i+1])
						//temp = item1
						temp:=list[name][i]
						//item1 = item2
						list[name][i]=list[name][i+1]
						//item2 = temp
						list[name][i+1]=temp
					}
				}
			}
		}
	}
}
//binary search function
//https://en.wikipedia.org/wiki/Binary_search_algorithm
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
	//fmt.Println("len : ",end)
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
