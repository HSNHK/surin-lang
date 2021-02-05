package syntax

import (
	"../../interpreter"
	"fmt"
	"strings"
)
//streql syntax
func StrEql(code string){
	StringA,StringB:=Tow(code)
	fmt.Println(interpreter.StrEql(strings.TrimSpace(StringA),strings.TrimSpace(StringB)))
}