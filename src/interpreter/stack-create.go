package interpreter

import "math/rand"

//CreateVariable function for create new variable
func CreateVariable(VariableName,_type string,stackMap *map[string][]interface{}){
	var stack map[string][]interface{}=*stackMap
	stack[VariableName]=[]interface{}{nil,nil,nil}
	stack[VariableName][0]=_type
	stack[VariableName][1]=nil
	stack[VariableName][2]=rand.Int()//is generate random id
}