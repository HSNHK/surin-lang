package interpreter

import (
	"fmt"
	"strings"
)

//system functions

func Print(data interface{}, dataType string){
	if dataType =="str"{
		fmt.Println(data.(string))
	}else if dataType =="int"{
		fmt.Println(data)
	}
}
func DataLen(value string)int{
	return len(value)
}

func Find(value1,value2 string)bool{
	if strings.Contains(value1,value2){
		return true
	}
	return false
}
