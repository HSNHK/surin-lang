package parser

//value type
const STRING ="str"
const INT ="int"
func MainParser(code string, VariablesStorageSpace, ListsStorageSpace *map[string][]interface{},line int){
	//Check the code namespace to navigate correctly to your class
	if IsValidNameSpace("List",code) {
		ListAnalyzer(code, ListsStorageSpace,line)
	}else {
		GlobalAnalyzer(code,VariablesStorageSpace,line)
	}
}