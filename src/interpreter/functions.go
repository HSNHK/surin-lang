package interpreter

import (
	"fmt"
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
func PushStack(VariableName string,Data interface{}, stackMap *map[string]interface{}){
	var stack map[string]interface{}=*stackMap
	stack[VariableName]=Data
}
func CreateVariable(VariableName string,stackMap *map[string]interface{}){
	var stack map[string]interface{}=*stackMap
	stack[VariableName]="nil"
}
func GetValue(VariableName string,stackMap *map[string]interface{})interface{}{
	var stack map[string]interface{}=*stackMap
	return stack[VariableName]
}
func ExistVariable(VariableName string,stackMap *map[string]interface{})bool{
	var stack map[string]interface{}=*stackMap
	for key:=range stack{
		if key==VariableName{
			return true
		}
	}
	return false
}
func DeleteVariable(VariableName string,stackMap *map[string]interface{})bool{
	var stack map[string]interface{}=*stackMap
	for key:=range stack{
		if key==VariableName{
			delete(stack,VariableName)
			return true
		}
	}
	return false
}
func AddToVariable(VariableName string,Data interface{}, stackMap *map[string]interface{}){
	var stack map[string]interface{}=*stackMap
	stack[VariableName]=fmt.Sprintf("%s+%s",stack[VariableName],Data)
}
