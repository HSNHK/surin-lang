package parser

import "regexp"

var(
	//all pattern | regex format
	//https://en.wikipedia.org/wiki/Regular_expression
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
		//cmp variable
		"cmp":"^[(cmp)|(\\scmp)|(cmp\\s)]+[(]+[a-z]+[\\,]+[a-z]+[)]$",
		//show variable type
		//type(variable-name)
		"var_type":"^[(type)|(\\stype)|(type\\s)]+[(]+[a-z]+[)]$",
		//show variable id
		//id(variable-name)
		"id":"^[(id)|(\\sid)|(id\\s)]+[(]+[a-z]+[)]$",
		//return variable value
		//ivar(variable-name)
		"ivar":"^[(ivar)|(\\sivar)|(ivar\\s)]+[(]+[a-z]+[)]$",
		//if style1
		//if(12,>=,114)?print("ok"):print("no")
		"if_s1":"^[(if)|(\\sif)|(if\\s)]+[(]+[\\d]+[,]+[(>)|(<)|(>=)|(<=)]+[,][0-9]+[)]+[?]+[(print)]+[(]+[\"].*[\"]+[)]+[:]+[(print)]+[(]+[\"].*[\"]+[)]$",
		//label
		//lb:name
		"label":"[(lb)|(\\slb)|(lb\\s)]+[(]+[a-z]+[)]$",
		//jmp
		//jmp(label-name)
		//"jmp":"[(jmptolab)|(\\sjmptolab)|(jmptolab\\s)]+[(]+[a-z]+[)]$",
	}
)

var List =map[string]string{
	//list
	//list(name:size)
	"list":"^[(list)|(\\slist)|(list\\s)]+[(]+[a-z]+[:]+[0-9]+[)]$",
	//list get value
	//list(name,index)
	"list-get":"^[(list)|(\\slist)|(list\\s)]+[(]+[a-z]+[,][0-9]+[)]$",
	//list push
	//list(name,index,value)
	"list-push":"^[(list)|(\\slist)|(list\\s)]+[(]+[a-z]+[,][0-9]+[,]+.*[)]$",
	//list show all item
	//list(name)
	"show-all-list-item":"^[(list)|(\\slist)|(list\\s)]+[(]+[a-z]+[)]$",
	//method
	//search to list
	//list.index(list-name,value)
	"list-search":"^[(list.search)|(\\slist.search)|(list.search\\s]+[(]+[a-z]+[,]+.*[)]$",
	//sort list
	//list.sort(list-name)
	"list-sort":"^[(list.sort)|(\\slist.sort)|(list.sort\\s]+[(]+[a-z]+[)]$",
	//delete list
	//list.delete(list-name)
	"list-delete":"^[(list.del)|(\\slist.del)|(list.del\\s]+[(]+[a-z]+[)]$",
}


//check match pattern with syntax
func IsValid(NameSpace,PTN,syntax string) bool {
	re, _ := regexp.MatchString(GetListPattern(NameSpace)[PTN], syntax)
	if re {
		return true
	}
	return false
}
//check exist pattern to namespace
func IsExistPattern(NameSpace,PTN string)bool{
	if _,ok:=GetListPattern(NameSpace)[PTN];ok{
		return true
	}
	return false
}
//return the number of all patterns
func PatternLen(NameSpace string)int{
	return len(GetListPattern(NameSpace))
}
//return all pattern key
func AllPatternKey(NameSpace string)string{
	var all string
	for key:=range GetListPattern(NameSpace){
		all=all+"\n"+key
	}
	return all
}