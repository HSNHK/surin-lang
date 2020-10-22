package list_syntax

import "../../../interpreter/list-structure"

func List(code string,ListMap *map[string][]interface{})  {
	list_structure.ListCreate(code[5:len(code)-1],ListMap)
}

