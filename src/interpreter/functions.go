package interpreter

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

//system functions
//Print function for print string or int or float or ...
func Print(data interface{}, dataType string){
	if dataType =="str"{
		fmt.Println(data.(string))
	}else if dataType =="int"{
		fmt.Println(data)
	}
}
//DataLen function for return string len
func DataLen(value string)int{
	return len(value)
}
//find function for find string2 to string1
func Find(value1,value2 string)bool{
	if strings.Contains(value1,value2){
		return true
	}
	return false
}
//PushStack function for push value to variable
func PushStack(VariableName,_type string,Data interface{}, stackMap *map[string][]interface{}){
	var stack map[string][]interface{}=*stackMap
	stack[VariableName][0]=_type
	stack[VariableName][1]=Data
}
//CreateVariable function for create new variable
func CreateVariable(VariableName,_type string,stackMap *map[string][]interface{}){
	var stack map[string][]interface{}=*stackMap
	stack[VariableName]=[]interface{}{nil,nil,nil}
	stack[VariableName][0]=_type
	stack[VariableName][1]=nil
	stack[VariableName][2]=rand.Int()
}
//GetValue function for return variable value
func GetValue(VariableName string,stackMap *map[string][]interface{})interface{}{
	var stack map[string][]interface{}=*stackMap
	return stack[VariableName][1]
}
//ExistVariable function for check variable exist
func ExistVariable(VariableName string,stackMap *map[string][]interface{})bool{
	var stack map[string][]interface{}=*stackMap
	for key:=range stack{
		if key==VariableName{
			return true
		}
	}
	return false
}
//DeleteVariable function for delete variable
func DeleteVariable(VariableName string,stackMap *map[string][]interface{})bool{
	var stack map[string][]interface{}=*stackMap
	for key:=range stack{
		if key==VariableName{
			//delete variable
			delete(stack,VariableName)
			return true
		}
	}
	return false
}
//AddToVariable function for add value to variable(last value + new value)
func AddToVariable(VariableName, _type string,Data interface{},stackMap *map[string][]interface{}) {
	var stack map[string][]interface{}=*stackMap
	//check variable type
	if _type=="str"{
		stack[VariableName][0]=_type
		//if variable type = str (str+str) hel+lo=hello
		stack[VariableName][1]= stack[VariableName][1].(string)+Data.(string)
	}else if _type=="int"{
		stack[VariableName][0]=_type
		//if variable type = int (last value + new value)
		//last value
		x,err:=strconv.Atoi(stack[VariableName][1].(string))
		//new value
		y,err:=strconv.Atoi(Data.(string))
		if err ==nil{
			stack[VariableName][1]= y + x
		}else{
			fmt.Println(err)
		}
	}
}
