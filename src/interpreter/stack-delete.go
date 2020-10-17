package interpreter

//DeleteVariable function for delete variable
func DeleteVariable(VariableName string,stackMap *map[string][]interface{})bool{
	var stack map[string][]interface{}=*stackMap
	for key:=range stack{
		if key==VariableName{
			//delete variable
			delete(stack,VariableName)
			return true//exist variable
		}
	}
	return false //not exist variable
}
