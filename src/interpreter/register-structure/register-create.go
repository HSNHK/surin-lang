package register_structure

import "math/rand"

//CreateVariable function for create new variable
func CreateVariable(VariableName,_type string, registerMap *map[string][]interface{}){
	var register map[string][]interface{} =*registerMap
	register[VariableName]=[]interface{}{nil,nil,nil}
	register[VariableName][0]=_type
	register[VariableName][1]=nil
	register[VariableName][2]=rand.Int() //is generate random id
}