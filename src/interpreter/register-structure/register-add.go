package register_structure

import (
	"fmt"
	"strconv"
)

//AddToVariable function for add value to variable(last value + new value)
func AddToVariable(VariableName, _type string,Data interface{}, registerMap *map[string][]interface{}) {
	var stack map[string][]interface{}=*registerMap
	//check variable type
	if _type=="str"{
		stack[VariableName][0]=_type
		//if variable type = str (str+str) hel+lo=hello
		stack[VariableName][1]= stack[VariableName][1].(string)+Data.(string)
	}else if _type=="int"{
		stack[VariableName][0]=_type
		//if variable type = int (last value + new value)
		//last value
		x,err:=strconv.Atoi(stack[VariableName][1].(string))
		//new value
		y,err:=strconv.Atoi(Data.(string))
		if err ==nil{
			//[1]=value index
			stack[VariableName][1]= y + x
		}else{
			//The value could not be converted to int
			fmt.Println(err)
		}
	}
}
