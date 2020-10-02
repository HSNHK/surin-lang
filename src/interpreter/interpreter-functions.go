package interpreter

import (
	"fmt"
)

//system functions

func Print(data interface{}, dataType string){
	if dataType =="str"{
		fmt.Println(data)
	}else if dataType =="int"{
		fmt.Println(data)
	}
}

