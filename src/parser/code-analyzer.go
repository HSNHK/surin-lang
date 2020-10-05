package parser

import (
	"../interpreter"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Core(code string){

	pattern:=map[string]string{
		"print":"^[(print)]+[(]+[\"].*[\"]+[)]$",
		"print_int":"^[(print)|(\\sprint)|(print\\s)]+[(][\\d]+[)]$",
		"logic":"^[(logic)|(\\slogic)|(logic\\s)]+[(]+[\\d]+[,]+[(>)|(<)|(>=)|(<=)]+[,][0-9]+[)]$",
		"math":"^[(math)|(\\smath)|(math\\s)]+[(]+[\\d]+[,]+(\\+|\\-|\\*|\\/)+[,]+[\\d]+[)]|[)]\\s$",
		"streql":"^[(streql)|(\\sstreql)|(streql\\s)]+[(]+[\"].*[\"]+[\\,]+[\"].*[\"]+[)]$",
		"len":"^[(len)|(\\slen)|(len\\s)]+[(]+[\"].*[\"]+[)]$",
	}

	if re,_:=regexp.MatchString(pattern["print"],code);re== true{
		interpreter.Print(code[7:len(code)-2], "str")

	}else if re,_:=regexp.MatchString(pattern["print_int"],code);re==true{
		interpreter.Print(code[7:len(code)-2], "int")

	}else if re,_:=regexp.MatchString(pattern["logic"],code);re== true{
		codeSpit :=strings.Split(code,",")
		v1,err1:=strconv.Atoi(codeSpit[0][6:])
		v2,err2:=strconv.Atoi(codeSpit[2][:len(codeSpit[2])-1])
		if err1==nil && err2==nil{
			fmt.Println(interpreter.Logic(codeSpit[1],v1,v2))
		}else{
			fmt.Println("logic format error simple :  logic(5,>=,3)")
		}
	}else if re,_:=regexp.MatchString(pattern["math"],code);re== true{
		codeSpit :=strings.Split(code,",")
		v1,err1:=strconv.Atoi(codeSpit[0][5:])
		v2,err2:=strconv.Atoi(codeSpit[2][:len(codeSpit[2])-1])
		if err1==nil && err2==nil{
			fmt.Println(interpreter.Math(codeSpit[1],v1,v2))
		}else{
			fmt.Println("logic format error simple : math(5,*,2)")
		}

	}else if re,_:=regexp.MatchString(pattern["streql"],code);re== true{
		codeSplit:=strings.Split(code,",")
		fmt.Println(interpreter.StrEql(
			strings.TrimSpace(codeSplit[1][1:len(codeSplit[1])-2]),
			strings.TrimSpace(codeSplit[1][1:len(codeSplit[1])-2])))

	}else if re,_:=regexp.MatchString(pattern["len"],code);re== true{
		fmt.Println(interpreter.DataLen(strings.TrimSpace(code[5:len(code)-2])))
	}
}
