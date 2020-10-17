package interpreter

import (
	"fmt"
	"strings"
	"time"
)

//system functions
//Print function for print string or int or float or ...
func Print(data interface{}, dataType string){
	if dataType =="str"{
		fmt.Println(data.(string))
	}else if dataType =="int"{
		fmt.Println(data)
	}
}
//DataLen function for return string len
func DataLen(value string)int{
	return len(value)
}
//find function for find string2 to string1
func Find(value1,value2 string)bool{
	if strings.Contains(value1,value2){
		return true
	}
	return false//not find
}
//time function
func Time() string{
	return fmt.Sprintf("%d:%d:%d",time.Now().Hour(),time.Now().Minute(),time.Now().Second())
}