package parser

//value type
const STRING ="str"
const INT ="int"
//The first code is checked to see what the namespace is and
//then sent to the related function for processing.
func MainParser(code string, VariablesStorageSpace, ListsStorageSpace *map[string][]interface{},line int){
	//Check the code namespace to navigate correctly to your class
	if IsValidNameSpace("List",code) {
		//List syntax are included in this namespace
		ListAnalyzer(code, ListsStorageSpace,line)
	}else {
		//Language-related syntax are included in this namespace
		GlobalAnalyzer(code,VariablesStorageSpace,line)
	}
}