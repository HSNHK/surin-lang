package parser

import (
	"../interpreter"
	"fmt"
	"strconv"
	"strings"
)

func Core(code string){

	if IsValid("print",code){
		interpreter.Print(code[7:len(code)-2], "str")

	}else if IsValid("print_int",code){
		interpreter.Print(code[7:len(code)-2], "int")

	}else if IsValid("logic",code){
		codeSpit :=strings.Split(code,",")
		v1,err1:=strconv.Atoi(codeSpit[0][6:])
		v2,err2:=strconv.Atoi(codeSpit[2][:len(codeSpit[2])-1])
		if err1==nil && err2==nil{
			fmt.Println(interpreter.Logic(codeSpit[1],v1,v2))
		}else{
			fmt.Println("logic format error simple :  logic(5,>=,3)")
		}
	}else if IsValid("math",code){
		codeSpit :=strings.Split(code,",")
		v1,err1:=strconv.Atoi(codeSpit[0][5:])
		v2,err2:=strconv.Atoi(codeSpit[2][:len(codeSpit[2])-1])
		if err1==nil && err2==nil{
			fmt.Println(interpreter.Math(codeSpit[1],v1,v2))
		}else{
			fmt.Println("logic format error simple : math(5,*,2)")
		}

	}else if IsValid("streql",code){
		codeSplit:=strings.Split(code,",")
		fmt.Println(interpreter.StrEql(
			strings.TrimSpace(codeSplit[0][8:len(codeSplit[0])-1]),
			strings.TrimSpace(codeSplit[1][1:len(codeSplit[1])-2])))

	}else if IsValid("len",code){
		fmt.Println(interpreter.DataLen(strings.TrimSpace(code[5:len(code)-2])))

	}else if IsValid("find",code){
		codeSplit:=strings.Split(code,",")
		fmt.Println(interpreter.Find(strings.TrimSpace(codeSplit[0][6:len(codeSplit[0])-1]),
			strings.TrimSpace(codeSplit[1][1:len(codeSplit[1])-2])))
	}
}
