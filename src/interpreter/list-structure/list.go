package list_structure

func ExistList(listName string,ListMap *map[string][]interface{}) bool{
	var list map[string][]interface{}=*ListMap
	for key:=range list{
		if listName==key{
			return true
		}
	}
	return false
}
