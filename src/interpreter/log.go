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
		fmt.Println(fmt.Sprintf("[%s]|%s|=>|%s|=> %s",infoFlag,Time(),location,log))
	}else if level==warning{
		fmt.Println(fmt.Sprintf("[%s]|%s|=>|%s|=> %s",warningFlag,Time(),location,log))
	}else if level==error {
		fmt.Println(fmt.Sprintf("[%s]|%s|=>|%s|=> %s",errorFlag,Time(),location,log))
	}
}