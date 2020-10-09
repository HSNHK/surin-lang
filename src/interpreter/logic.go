package interpreter

//StrEql function for cmp(string,string)
func StrEql(str1,str2 string)bool{
	if str1==str2{
		return true
	}
	//not math strings
	return false
}
//logic function(=,>,<,>=,<=)
func Logic(LogicType string,value1,value2 int)bool {
	switch LogicType {
	case "=":
		//check value1 = value2
		if value1 == value2 {
			return true
		} else {
			return false
		}
		break
		//check value1 > value2
	case ">":
		if value1 > value2 {
			return true
		} else {
			return false
		}
		break
		//check value1 < value2
	case "<":
		if value1 < value2 {
			return true
		} else {
			return false
		}
		break
		//check value1 >= value2
	case ">=":
		if value1 >= value2 {
			return true
		} else {
			return false
		}
		break
		//check value1 <= value2
	case "<=":
		if value1 <= value2 {
			return true
		} else {
			return false
		}
		break
	}
	//not exist operation
	return false
}
