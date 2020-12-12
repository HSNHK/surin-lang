package parser

import "regexp"


var Namespace=map[string][]interface{}{
	"Global":{pattern},
	"List":{"^[list]+.*$",List},
}

//get name space pattern with name
func GetNameSpacePattern(name string)string{
	return Namespace[name][0].(string)
}
//get name space list pattern
func GetListPattern(name string)map[string]string{
	//all pattern
	if name!="Global"{
		return Namespace[name][1].(map[string]string)
	}
	//global pattern
	return Namespace[name][0].(map[string]string)
}
//check exits pattern
func IsNameSpace(name string)bool{
	if Namespace[name]!=nil{
		return true
	}
	return false
}
//check code name space valid
func IsValidNameSpace(namespace,code string)bool{
	re, _ := regexp.MatchString(GetNameSpacePattern(namespace), code)
	if re == true {
		return true
	}
	return false
}