package main

import (
	"./parser"
	"bufio"
	"fmt"
	"log"
	"os"
)
//stack map
//[variable-name]:[variable-type(int,str),variable-value,ID]
var stack=map[string][]interface{}{
	"first":{"type","value","id"},
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
		for fileScan.Scan(){
			//check commented
			if fileScan.Text()[0]== '#' {
				continue
			}
			//&stack=send stack address or send reference
			parser.Core(fileScan.Text(),&stack)
		}
		//file open error
		if err:= fileScan.Err();err!=nil{
			log.Fatal(fmt.Sprintf("[%s]",err))
		}

	}else {
		//constant syntax
		command:=map[string]string{
			"exit":"enter exit() for close in interpreter",
			"hello":"hello programmer welcome to surin interpreter",
			"help":"pales view README",
			"true":"true",
			"false":"false",
		}
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
				//&stack=send stack address or send reference
			}else{
				parser.Core(userCommand,&stack)
			}
		}
	}
}
