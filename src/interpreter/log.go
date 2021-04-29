package interpreter

import (
	"fmt"
)
//log level and flag
const(
	info=1
	infoFlag="I!"
	warning=2
	warningFlag="W/"
	error=3
	errorFlag="E*"
)
//log function
func Log(log,location string,level int)  {
	if level==info{
		fmt.Printf("[%s]|%s|=>|%s|=> %s\n",infoFlag,Time(),location,log)
	}else if level==warning{
		fmt.Printf("[%s]|%s|=>|%s|=> %s\n",warningFlag,Time(),location,log)
	}else if level==error {
		fmt.Printf("[%s]|%s|=>|%s|=> %s\n",errorFlag,Time(),location,log)
	}
}
//not found syntax
func NotFound(code string,line int){
	fmt.Println(
		fmt.Sprintf("+=+=+=+=+=+=+=+=+=+=+=+=+=+=+\n"+
				"|*| syntax not found\n" +
				"|*| line :[%d]\n"+
				"|*| code syntax :[%s]\n"+
		"+=+=+=+=+=+=+=+=+=+=+=+=+=+=+",line,code))
}
