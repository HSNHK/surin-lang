package parser

import "regexp"

var(
	//all pattern | regex format
	pattern=map[string]string{
		//print pattern
		//print("hello")
		"print":"^[(print)]+[(]+[\"].*[\"]+[)]$",
		//print(123456)
		"print_int":"^[(print)|(\\sprint)|(print\\s)]+[(][\\d]+[)]$",
		//logic pattern
		//logic(12,>=,12)
		"logic":"^[(logic)|(\\slogic)|(logic\\s)]+[(]+[\\d]+[,]+[(>)|(<)|(>=)|(<=)]+[,][0-9]+[)]$",
		//math pattern
		//math(12,**,2)
		"math":"^[(math)|(\\smath)|(math\\s)]+[(]+[\\d]+[,]+(\\+|\\-|\\*|\\/)+[,]+[\\d]+[)]|[)]\\s$",
		//string eql pattern
		//streql("hello","hello")
		"streql":"^[(streql)|(\\sstreql)|(streql\\s)]+[(]+[\"].*[\"]+[\\,]+[\"].*[\"]+[)]$",
		//data len pattern
		//len("hello im surin lang")
		"len":"^[(len)|(\\slen)|(len\\s)]+[(]+[\"].*[\"]+[)]$",
		//find to string pattern
		//find("hello im surin lang","surin")
		"find":"^[(find)|(\\sfind)|(find\\s)]+[(]+[\"].*[\"]+[\\,]+[\"].*[\"]+[)]$",
		//creat variable pattern
		//var(variable-name,"string value")
		"var_str":"^[(var)|(\\svar)|(var\\s)]+[(]+[a-z]+[\\,]+[str]+[)]$",
		//var(variable-name,int)
		"var_int":"^[(var)|(\\svar)|(var\\s)]+[(]+[a-z]+[\\,]+[int]+[)]$",
		//push to variable pattern
		//push(variable-name,"string value")
		"push_str":"^[(push)|(\\spush)|(push\\s)]+[(]+[\"].*[\"]+[\\,]+[\"].*[\"]+[)]$",
		//push(variable-name,int value)
		"push_int":"^[(push)|(\\spush)|(push\\s)]+[(]+[\"].*[\"]+[\\,]+[0-9]+[)]$",
		//pop variable value pattern
		//pop(variable-name)
		"pop":"^[(pop)|(\\spop)|(pop\\s)]+[(]+[a-z]+[)]$",
		//remove variable pattern
		//rm(variable-name)
		"rm":"^[(rm)|(\\srm)|(rm\\s)]+[(]+[a-z]+[)]$",
		//add to variable value pattern
		//add(variable-name,int value)
		"add_int":"^[(add)|(\\sadd)|(add\\s)]+[(]+[\"].*[\"]+[\\,]+[0-9]+[)]$",
		//add(variable-name,"string value")
		"add_str":"^[(add)|(\\sadd)|(add\\s)]+[(]+[\"].*[\"]+[\\,]+[\"].*[\"]+[)]$",
		//show variable type
		//type(variable-name)
		"type":"^[(type)|(\\stype)|(type\\s)]+[(]+[a-z]+[)]$",
		//show variable id
		//id(variable-name)
		"id":"^[(id)|(\\sid)|(id\\s)]+[(]+[a-z]+[)]$",
		//return variable value
		//ivar(variable-name)
		"ivar":"^[(ivar)|(\\sivar)|(ivar\\s)]+[(]+[a-z]+[)]$",
		//if style1
		//if(12,>=,114)?print("ok"):print("no")
		"if_s1":"^[(if)|(\\sif)|(if\\s)]+[(]+[\\d]+[,]+[(>)|(<)|(>=)|(<=)]+[,][0-9]+[)]+[?]+[(print)]+[(]+[\"].*[\"]+[)]+[:]+[(print)]+[(]+[\"].*[\"]+[)]$",
	}
)
//check match pattern with syntax
func IsValid(PTN,syntax string) bool {
	re, _ := regexp.MatchString(pattern[PTN], syntax)
	if re == true {
		return true
	}
	return false
}
//check exist pattern
func IsExistPattern(PTN string)bool{
	for i:=range pattern{
		if i==PTN{
			return true
		}
	}
	return false
}
//return the number of all patterns
func PatternLen()int{
	return  len(pattern)
}
//return all pattern key
func AllPatternKey()string{
	var all string
	for i:=range pattern{
		all=all+"\n"+i
	}
	return all
}