package interpreter

//Push register function for push value to variable
func PushStack(VariableName,_type string,Data interface{}, registerMap *map[string][]interface{}){
	var stack map[string][]interface{}=*registerMap
	stack[VariableName][0]=_type
	stack[VariableName][1]=Data
}
