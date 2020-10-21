package register_structure


//GetValue function for return variable value
func GetValue(VariableName string, registerMap *map[string][]interface{})interface{}{
	var stack map[string][]interface{}=*registerMap
	return stack[VariableName][1]//[1]=value index
}
