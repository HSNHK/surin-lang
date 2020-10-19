package interpreter


//return a variable type
func Type(VariableName string, stackMap *map[string][]interface{}) interface{} {
	var stack map[string][]interface{}=*stackMap
	for key:=range stack{
		if key==VariableName{
			return stack[VariableName][0]//type index
		}
	}
	return nil //not exist variable
}
//return a variable id
//create id in CreateVariable function
func Id(VariableName string, stackMap *map[string][]interface{}) interface{} {
	var stack map[string][]interface{}=*stackMap
	for key:=range stack{
		if key==VariableName{
			return stack[VariableName][2]//id index
		}
	}
	return nil //not exist variable
}

//ExistVariable function for check variable exist
func ExistVariable(VariableName string,stackMap *map[string][]interface{})bool{
	var stack map[string][]interface{}=*stackMap
	for key:=range stack{
		if key==VariableName{
			return true //exist variable
		}
	}
	return false //not exist variable
}
//cmp function
func Cmp(var1,var2 string,stackMap *map[string][]interface{}) bool {
	if GetValue(var1,stackMap)==GetValue(var2,stackMap){
		return true
	}else{
		return false
	}
}