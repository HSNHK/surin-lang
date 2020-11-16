package main

import (
	"./interpreter"
	"./parser"
	"bufio"
	"fmt"
	"log"
	"os"
)
//registers map
//[variable-name]:[variable-type(int,str),variable-value,ID]
var registers =map[string][]interface{}{
	"first":{"type","value","id"},
}
//list map
//[list-name]:[item]
var list =map[string] []interface{}{
	"name":{1,2,3,4,5,6},
}
//list code
//[code-line-1,code-line-2,...]
var CodeList=map[int]string{
	//1:"print(3)",
}
var head =0
//constant syntax
var command=map[string]interface{}{
		"exit":"enter exit() for close in interpreter",
		"hello":"hello programmer welcome to surin interpreter",
		"help":"pales view README",
		"true":"true",
		"false":"false",
		"time":interpreter.Time(),
		"info":interpreter.Info(),
}

func main(){
	input:=os.Args
	//check input = surin.exe file-path
	if len(input)>1{
		file,err:=os.Open(input[1])
		if err!=nil{
			log.Fatal(fmt.Sprintf("[%s]"," not open source code"))
		}
		//open file
		fileScan :=bufio.NewScanner(file)
		/*
		for fileScan.Scan(){
			//check commented
			if fileScan.Text()[0]== '#' {
				continue
			}
			for key,value:=range command{
				if fileScan.Text()==key{
					fmt.Println(value)
					continue
				}
			}
			//&registers=send registers address or send reference
			parser.Core(fileScan.Text(),&registers,&list)
		}
		 */
		for fileScan.Scan(){
			if fileScan.Text()[0]== '#' {
				continue
			}
			CodeList[head]=fileScan.Text()
			head++
		}
		for head:=0;head<=len(CodeList);head++ {
			for key,value:=range command{
				if CodeList[head]==key{
					fmt.Println(value)
					continue
				}
			}
			parser.Core(CodeList[head], &registers, &list)
		}
		//file open error
		if err:= fileScan.Err();err!=nil{
			log.Fatal(fmt.Sprintf("[%s]",err))
		}

	}else {

		//baner
		fmt.Println("   _____            _          __")
		fmt.Println("  / ___/__  _______(_)___     / /   ____ _____  ____ _")
		fmt.Println("  \\__ \\/ / / / ___/ / __ \\   / /   / __ `/ __ \\/ __ `/")
		fmt.Println(" ___/ / /_/ / /  / / / / /  / /___/ /_/ / / / / /_/ /")
		fmt.Println("/____/\\__,_/_/  /_/_/ /_/  /_____/\\__,_/_/ /_/\\__, /")
		fmt.Println("                                             /____/")
		fmt.Println("Hello Welcome to surin interpreter\n")

		//interpreter shell mode
		var userCommand string
		for true {

			fmt.Print(">>~#>")
			fmt.Scan(&userCommand)

			for key,value:=range command{
				if userCommand==key{
					fmt.Println(value)
					break
				}
			}
			if userCommand =="\n" || userCommand==" " {
				continue
			} else if userCommand=="exit()"{
				break
				//&registers=send registers address or send reference
			}else{
				parser.Core(userCommand,&registers,&list)
				CodeList[head]=userCommand
				head++
			}
		}
	}
}
