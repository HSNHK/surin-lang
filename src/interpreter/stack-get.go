package interpreter


//GetValue function for return variable value
func GetValue(VariableName string,stackMap *map[string][]interface{})interface{}{
	var stack map[string][]interface{}=*stackMap
	return stack[VariableName][1]//[1]=value index
}
