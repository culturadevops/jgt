package jjson

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/culturadevops/jgt/jio"
	"github.com/culturadevops/jgt/jlog"
)

type Jjson struct {
	Log *jlog.Jlog
}

func (i *Jjson) CreateFileWithStruct(fileName, jsonestruct string) {
	if jsonestruct == "" {
		jsonestruct = `{"param1":"dat0","param2":"dat1"}`
	}
	jio.CreateFile(fileName, jsonestruct)
}

/*
lee un archivo sin estructuras
para usar debes hacer esto

GetJsonFromFile("miarchivo.json")
*/
func (i *Jjson) GetJsonFromFile(fileConfigName string) map[string]interface{} {

	jsonFile, err := os.Open(fileConfigName)
	if err != nil {
		if err.Error() == "open "+fileConfigName+": no such file or directory" {
			i.Log.IsFatal(err)
		}
	}
	defer jsonFile.Close()
	var result map[string]interface{}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &result)
	return result
}

/*
lee un archivo buscando una estructura y la devuelve
para usar debes hacer esto
type FilesStruct struct {
	Name  string
}
var WithStruct []FilesStruct
GetJsonFileWithStruct("miarchivo.json", &WithStruct)
*/
func (i *Jjson) GetJsonFileWithStruct(jsonFileName string, WithStruct interface{}) {
	jsonFile, err := os.Open(jsonFileName)
	i.Log.IsFatal(err)
	defer jsonFile.Close()
	byteValue, err1 := ioutil.ReadAll(jsonFile)
	i.Log.IsFatal(err1)
	i.Log.IsFatal(json.Unmarshal([]byte(byteValue), WithStruct))
}
