package syntax

import (
	"../../interpreter"
	"fmt"
	"strings"
)
//find function syntax
func Find(code string)  {
	ValueA,ValueB:=Tow(code)
	fmt.Println(interpreter.Find(strings.TrimSpace(ValueA), strings.TrimSpace(ValueB)))
}
