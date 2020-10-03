package parser

import (
	"../interpreter"
	"fmt"
	"strconv"
	"strings"
)

func Core(code string){
	if code[:5]=="print" {
		if code[6] == '"' {
			interpreter.Print(code[7:len(code)-2], "str")
		} else {
			interpreter.Print(code[6:len(code)-1], "int")
		}

	}else if code[:5]=="logic"{
		codeSpit :=strings.Split(code,",")
		v1,err1:=strconv.Atoi(codeSpit[0][6:])
		v2,err2:=strconv.Atoi(codeSpit[2][:len(codeSpit[2])-1])
		if err1==nil && err2==nil{
			fmt.Println(interpreter.Logic(codeSpit[1],v1,v2))
		}else{
			fmt.Println("logic format error simple :  logic(5,>=,3)")
		}

	}else if code[:4]=="math"{
		codeSpit :=strings.Split(code,",")
		v1,err1:=strconv.Atoi(codeSpit[0][5:])
		v2,err2:=strconv.Atoi(codeSpit[2][:len(codeSpit[2])-1])
		if err1==nil && err2==nil{
			fmt.Println(interpreter.Math(codeSpit[1],v1,v2))
		}else{
			fmt.Println("logic format error simple : math(5,*,2)")
		}

	}else if code[:6]=="streql"{
		codeSplit:=strings.Split(code,",")
		fmt.Println(interpreter.StrEql(
			strings.TrimSpace(codeSplit[1][1:len(codeSplit[1])-2]),
			strings.TrimSpace(codeSplit[1][1:len(codeSplit[1])-2])))

	}else if code[:3]=="len"{
		fmt.Println(interpreter.DataLen(strings.TrimSpace(code[5:len(code)-2])))
	}
}
