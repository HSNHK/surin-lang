package register_structure


//return a variable type
func Type(VariableName string, registerMap *map[string][]interface{}) interface{} {
	var register map[string][]interface{} =*registerMap
	for key:=range register {
		if key==VariableName{
			return register[VariableName][0] //type index
		}
	}
	return nil //not exist variable
}
//return a variable id
//create id in CreateVariable function
func Id(VariableName string, registerMap *map[string][]interface{}) interface{} {
	var register map[string][]interface{} =*registerMap
	for key:=range register {
		if key==VariableName{
			return register[VariableName][2] //id index
		}
	}
	return nil //not exist variable
}

//ExistVariable function for check variable exist
func ExistVariable(VariableName string, registerMap *map[string][]interface{})bool{
	var register map[string][]interface{} =*registerMap
	for key:=range register {
		if key==VariableName{
			return true //exist variable
		}
	}
	return false //not exist variable
}
//cmp function
func Cmp(var1,var2 string, registerMap *map[string][]interface{}) bool {
	if GetValue(var1, registerMap)==GetValue(var2, registerMap){
		return true
	}else{
		return false
	}
}