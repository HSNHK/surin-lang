package parser

import (
	"./syntax/list-syntax"
)


func ListAnalyzer(code string,list *map[string][]interface{}) {
	//list(name:size)
	if IsValid("List","list", code) {
		list_syntax.List(code, list)

		//list(name,index,value)
	} else if IsValid("List","list-push", code) {
		list_syntax.ListPush(code, list)

		//list(name,index)
	} else if IsValid("List","list-get", code) {
		list_syntax.ListGet(code, list)

		//list(list-name)
	} else if IsValid("List","show-all-list-item", code) {
		list_syntax.ShowList(code, list)

		//list(list-name,item)
	} else if IsValid("List","list-search", code) {
		list_syntax.Search(code, list)

		//list(list-name)
	} else if IsValid("List","list-sort", code) {
		list_syntax.Sort(code, list)

		//list(list-name)
	} else if IsValid("List","list-delete", code) {
		list_syntax.Delete(code, list)
	}
}
