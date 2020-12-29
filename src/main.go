package main

import (
	"./interpreter"
	"./parser"
	"bufio"
	"fmt"
	"log"
	"os"
)

//VariablesStorageSpace map
//User-generated variables are stored in this map
//In such a way that the name of the map key variable
//and the type, value and value identifier of this map
//id is a random identifier
//and represents the variable identifier in memory
var VariablesStorageSpace =map[string][]interface{}{
	//[variable-name]:[variable-type(int,str),variable-value,ID]
	"first":{"type","value","id"},
}

//list map
//Arrays defined by the programmer are stored here
//This is a list of objects
//and different types of data can be stored in one list
var ListsStorageSpace =map[string] []interface{}{
	//[list-name]:[item]
	"name":{1,2,3,4,5,6},
}
//list code
//The code fetched in the name.sur file is stored here
//The reason for this is the ability to move lines of code
//Suppose we implement the jmp operation in this interpreter language.
//To do this, we need to have the previous line numbers so
//that we can move between the lines.
//[code-line-1,code-line-2,...]
var SourceCodeList =map[int]string{
	//1:"print(3)",
}
//Head is the line pointer and
//the value of this variable represents the current line number
var Head =0
//constant syntax
//Fixed commands in the interpreter can be added to this map
var ConstantCommand=map[string]interface{}{
		//Help to exit the interpreter
		"exit":"enter exit() for close in interpreter",
		//hello :)
		//This command is also available in Lisp
		"hello":"hello programmer welcome to surin interpreter",
		//help :)
		"help":"pales view README",
		"true":"true",
		"false":"false",
		//get system time
		"time":interpreter.Time(),
		//get system information
		"info":interpreter.Info(),
}

func main(){
	VariableArguments :=os.Args
	//Check the interpreter VariableArguments
	//check VariableArguments surin file-path
	if len(VariableArguments)>1{
		file,err:=os.Open(VariableArguments[1])
		if err!=nil{
			log.Fatal(fmt.Sprintf("[%s]"," not open source code"))
		}
		//open file
		SourceCodeScan :=bufio.NewScanner(file)
		for SourceCodeScan.Scan(){
			//Check that this line of code is in comment mode,
			//or if it is, this line will not execute.
			if SourceCodeScan.Text()[0]== '#' {
				continue
			}
			SourceCodeList[Head]= SourceCodeScan.Text()
			Head++
		}
		for head:=0;head<len(SourceCodeList);head++ {
			if value,ok:=ConstantCommand[SourceCodeList[head]];ok{
				fmt.Println(value)
				continue
			}
			parser.MainParser(SourceCodeList[head], &VariablesStorageSpace, &ListsStorageSpace,head)
		}
		//file open error
		if err:= SourceCodeScan.Err();err!=nil{
			log.Fatal(fmt.Sprintf("[%s]",err))
		}

	}else {
		//interpreter mode
		//baner
		fmt.Println("   _____            _          __")
		fmt.Println("  / ___/__  _______(_)___     / /   ____ _____  ____ _")
		fmt.Println("  \\__ \\/ / / / ___/ / __ \\   / /   / __ `/ __ \\/ __ `/")
		fmt.Println(" ___/ / /_/ / /  / / / / /  / /___/ /_/ / / / / /_/ /")
		fmt.Println("/____/\\__,_/_/  /_/_/ /_/  /_____/\\__,_/_/ /_/\\__, /")
		fmt.Println("                                             /____/")
		fmt.Println("Hello Welcome to surin interpreter\n")
		//interpreter shell mode
		var UserCommand string
		for true {
			fmt.Print(">>~#>")
			fmt.Scan(&UserCommand)
			for key,value:=range ConstantCommand{
				if UserCommand ==key{
					fmt.Println(value)
					break
				}
			}
			if UserCommand =="\n" || UserCommand ==" " {
				continue
			} else if UserCommand =="exit()"{
				break
				//&VariablesStorageSpace=send VariablesStorageSpace address or send reference
			}else{
				parser.MainParser(UserCommand,&VariablesStorageSpace,&ListsStorageSpace,0)
				SourceCodeList[Head]= UserCommand
				Head++
			}
		}
	}
}
