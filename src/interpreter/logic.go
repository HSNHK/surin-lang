package interpreter


func StrEql(str1,str2 string)bool{
	if str1==str2{
		return true
	}
	return false
}
func Logic(LogicType string,value1,value2 int)bool {
	switch LogicType {
	case "=":
		if value1 == value2 {
			return true
		} else {
			return false
		}
		break
	case ">":
		if value1 > value2 {
			return true
		} else {
			return false
		}
		break
	case "<":
		if value1 < value2 {
			return true
		} else {
			return false
		}
		break
	case ">=":
		if value1 >= value2 {
			return true
		} else {
			return false
		}
		break
	case "<=":
		if value1 <= value2 {
			return true
		} else {
			return false
		}
		break
	}
	return false
}
