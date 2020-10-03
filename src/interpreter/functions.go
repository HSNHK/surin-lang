package interpreter

import (
	"fmt"
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


