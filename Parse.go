package config
//create by libiao 20190107
//support json configFile
//sample like config.Parse(*s interface{},filename string) ,s is a struct


import (
	"strings"
	"os"
	"io/ioutil"
	"path"
	"fmt"
	"gopkg.in/yaml.v2"
)

type Bytes = []byte

//将yaml 文件转为换struct

type HandleFuncParse func(b Bytes, v interface{}) error
func (p HandleFuncParse) StartParse(b Bytes, v interface{}){
	p(b,v)
}



func  Yaml(b Bytes, v interface{}) error{
	return yaml.Unmarshal(b,v)
}


func Parse(file string,config interface{}) (err error){

	//get the context  of file , bytes
	fileWithDirectory := strings.Replace(file, "~", os.Getenv("HOME"), -1)
	b,err:=FileToBytes(fileWithDirectory)
	if err != nil{
		return err
	}

	// check the file suffix ,if .json then go to jsonformat
	// target: from json to struct
	suffix:=path.Ext(fileWithDirectory)
	switch suffix{
	case ".yaml":
		err=Yaml(b,config)
	}
	return err
}

// get the file context and return []byte
func FileToBytes(s string) (Bytes, error) {
	fmt.Println(s)
	fmt.Println(ioutil.ReadFile(s))
	return ioutil.ReadFile(s)
}






