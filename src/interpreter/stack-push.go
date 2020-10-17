package interpreter

//PushStack function for push value to variable
func PushStack(VariableName,_type string,Data interface{}, stackMap *map[string][]interface{}){
	var stack map[string][]interface{}=*stackMap
	stack[VariableName][0]=_type
	stack[VariableName][1]=Data
}
