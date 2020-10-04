package main

import (
	"./parser"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main(){
	input:=os.Args
	if len(input)>1{
		file,err:=os.Open(input[1])
		if err!=nil{
			log.Fatal(fmt.Sprintf("[%s]"," not open source code"))
		}
		fileScan :=bufio.NewScanner(file)
		for fileScan.Scan(){
			parser.Core(fileScan.Text())
		}
		if err:= fileScan.Err();err!=nil{
			log.Fatal(fmt.Sprintf("[%s]",err))
		}

	}else {
		command:=map[string]string{
			"exit":"enter exit() for close in interpreter",
			"hello":"hello programmer welcome to surin interpreter",
			"help":"pales view README",

		}

		fmt.Println("   _____            _          __")
		fmt.Println("  / ___/__  _______(_)___     / /   ____ _____  ____ _")
		fmt.Println("  \\__ \\/ / / / ___/ / __ \\   / /   / __ `/ __ \\/ __ `/")
		fmt.Println(" ___/ / /_/ / /  / / / / /  / /___/ /_/ / / / / /_/ /")
		fmt.Println("/____/\\__,_/_/  /_/_/ /_/  /_____/\\__,_/_/ /_/\\__, /")
		fmt.Println("                                             /____/")
		fmt.Println("Hello Welcome to surin interpreter\n")


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
			if userCommand=="exit()"{
				break
			}else{
				parser.Core(userCommand)
			}
		}
	}
}
