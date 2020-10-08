package interpreter

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

//system functions

func Print(data interface{}, dataType string){
	if dataType =="str"{
		fmt.Println(data.(string))
	}else if dataType =="int"{
		fmt.Println(data)
	}
}
func DataLen(value string)int{
	return len(value)
}
func Find(value1,value2 string)bool{
	if strings.Contains(value1,value2){
		return true
	}
	return false
}
func PushStack(VariableName,_type string,Data interface{}, stackMap *map[string][]interface{}){
	var stack map[string][]interface{}=*stackMap
	stack[VariableName][0]=_type
	stack[VariableName][1]=Data
}
func CreateVariable(VariableName,_type string,stackMap *map[string][]interface{}){
	var stack map[string][]interface{}=*stackMap
	stack[VariableName]=[]interface{}{nil,nil,nil}
	stack[VariableName][0]=_type
	stack[VariableName][1]=nil
	stack[VariableName][2]=rand.Int()
}
func GetValue(VariableName string,stackMap *map[string][]interface{})interface{}{
	var stack map[string][]interface{}=*stackMap
	return stack[VariableName][1]
}
func ExistVariable(VariableName string,stackMap *map[string][]interface{})bool{
	var stack map[string][]interface{}=*stackMap
	for key:=range stack{
		if key==VariableName{
			return true
		}
	}
	return false
}
func DeleteVariable(VariableName string,stackMap *map[string][]interface{})bool{
	var stack map[string][]interface{}=*stackMap
	for key:=range stack{
		if key==VariableName{
			delete(stack,VariableName)
			return true
		}
	}
	return false
}
func AddToVariable(VariableName, _type string,Data interface{},stackMap *map[string][]interface{}) {
	var stack map[string][]interface{}=*stackMap
	if _type=="str"{
		stack[VariableName][0]=_type
		stack[VariableName][1]= stack[VariableName][1].(string)+Data.(string)
	}else if _type=="int"{
		stack[VariableName][0]=_type
		x,err:=strconv.Atoi(stack[VariableName][1].(string))
		y,err:=strconv.Atoi(Data.(string))
		if err ==nil{
			stack[VariableName][1]= y + x
		}else{
			fmt.Println(err)
		}
	}
}
