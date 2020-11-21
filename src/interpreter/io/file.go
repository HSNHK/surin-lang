package io
import (
	"os"
)

func Create(path string)bool{
	_,err :=os.Create(path)
	if err!=nil{
		return false
	}
	return true
}
func Delete(path string)bool{
	err:=os.Remove(path)
	if err!=nil{
		return  false
	}
	return true
}
func Remove(path ,new string)bool{
	err:=os.Rename(path,new)
	if err!=nil{
		return false
	}
	return true
}
func Open(path string)(interface{},error) {
	file,err:=os.Open(path)
	if err!=nil {
		return nil, err
	}
	b:=make([]byte,5)
	return  file.Read(b)
}