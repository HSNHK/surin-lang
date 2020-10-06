package parser

import "regexp"

var(
	pattern=map[string]string{
		"print":"^[(print)]+[(]+[\"].*[\"]+[)]$",
		"print_int":"^[(print)|(\\sprint)|(print\\s)]+[(][\\d]+[)]$",
		"logic":"^[(logic)|(\\slogic)|(logic\\s)]+[(]+[\\d]+[,]+[(>)|(<)|(>=)|(<=)]+[,][0-9]+[)]$",
		"math":"^[(math)|(\\smath)|(math\\s)]+[(]+[\\d]+[,]+(\\+|\\-|\\*|\\/)+[,]+[\\d]+[)]|[)]\\s$",
		"streql":"^[(streql)|(\\sstreql)|(streql\\s)]+[(]+[\"].*[\"]+[\\,]+[\"].*[\"]+[)]$",
		"len":"^[(len)|(\\slen)|(len\\s)]+[(]+[\"].*[\"]+[)]$",
		"find":"^[(find)|(\\sfind)|(find\\s)]+[(]+[\"].*[\"]+[\\,]+[\"].*[\"]+[)]$",
	}
)

func IsValid(PTN,syntax string) bool {
	re, _ := regexp.MatchString(pattern[PTN], syntax)
	if re == true {
		return true
	}
	return false
}

func IsExistPattern(PTN string)bool{
	for i:=range pattern{
		if i==PTN{
			return true
		}
	}
	return false
}

func PatternLen()int{
	return  len(pattern)

}

func AllPatternKey()string{
	var all string;
	for i:=range pattern{
		all=all+"\n"+i
	}
	return all
}