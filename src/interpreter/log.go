package interpreter

import (
	"fmt"
	"time"
)
const(
	info=1
	infoFlag="I!"
	warning=2
	warningFlag="W/"
	error=3
	errorFlag="E*"
)
func Log(log,location string,level int)  {
	if level==info{
		fmt.Println(fmt.Sprintf("[%s][%s] => [%s] => [%s]",infoFlag,time.Now(),location,log))
	}else if level==warning{
		fmt.Println(fmt.Sprintf("[%s][%s] => [%s] => [%s]",warningFlag,time.Now(),location,log))
	}else if level==error {
		fmt.Println(fmt.Sprintf("[%s][%s] => [%s] => [%s]",errorFlag,time.Now(),location,log))
	}
}