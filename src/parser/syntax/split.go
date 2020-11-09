package syntax

import "strings"

//function(input)
func One(code string)string{
	//function([input)]
	codesplit:=strings.Split(code,"(")
	//function([input])
	return codesplit[1][:len(codesplit)-1]
}
//function(input-x,input-y)
func Tow(code string)(string,string){
	//function([input-x,input-y)]
	codesplit1 :=strings.Split(code,"(")
	//function(input-x[,]input-y)
	codesplit2 :=strings.Split(codesplit1[1],",")
	//function([input-x],[input-y])
	return codesplit2[0], codesplit2[1][:len(codesplit2[1])-1]
}
//function(input-x,input-y,input-z)
func Three(code string)(string,string,string){
	//function([input-x,input-y,input-z)]
	codesplit1 :=strings.Split(code,"(")
	//function(input-x[,]input-y[,]input-z)
	codesplit2 :=strings.Split(codesplit1[1],",")
	//function([input-x],[input-y],[input-z])
	return codesplit2[0],codesplit2[1],codesplit2[2][:len(codesplit2)-1]
}
//function(type:Initialise)
func Initialise(code string)(string,string){
	//function([type:Initialise)]
	codesplit1 :=strings.Split(code,"(")
	//function(type[:]Initialise)
	codesplit2 :=strings.Split(codesplit1[1],":")
	//function([type]:[Initialise])
	return codesplit2[0],codesplit2[1][:len(codesplit2[1])-1]
}