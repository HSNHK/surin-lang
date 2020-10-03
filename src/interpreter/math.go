package interpreter

func power(value1,value2 int) int{
	if value2==1{
		return value1
	}
	return value1*power(value1,value2-1)
}

func fact(n int)int{
	if n==1{
		return n
	}
	return n*fact(n-1)
}

func Math(mathType string,value1,value2 int) int{
	switch mathType {
	case "+":
		return value1+value2
	case "-":
		return value1-value2
	case "*":
		return value1*value2
	case "/":
		return value1/value2
	case "**":
		return power(value1,value2)
	case "fact":
		return fact(value1)
	}
	return 0
}
