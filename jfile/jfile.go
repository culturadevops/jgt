package jfile

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/culturadevops/jgt/jlog"
)

type Jfile struct {
	Log               *jlog.Jlog
	OrigHomedir       string
	OrigFolder        string
	DestinationDir    string
	DestinationFolder string
	Map               map[string]string
}

/*
Agrega datos para usar las variables de folder home y origen
estas variables son usadas en las siguientes funciones
GetTempleteName
GetDestinationFileName
*/
func (i *Jfile) SetOriginData(OrigHomedir string, OrigFolder string) {
	i.OrigHomedir = OrigHomedir
	i.OrigFolder = OrigFolder
}

/* Agrega datos a las variables de destino para las funciones de crear template y archivos
 */
func (i *Jfile) SetDestinationData(DestinationDir string, DestinationFolder string) {
	i.DestinationDir = DestinationDir
	i.DestinationFolder = DestinationFolder
}

/*
Genera una ruta de una plantilla usando origin home y origin folder  de Jfile
*/
func (i *Jfile) GetOriginFile(OriginFileName string) string {
	return i.GetTempleteName(i.OrigFolder, OriginFileName)
}

/*
Genera ruta de destino usando destinohome y destion folder de el objeto jfile
*/
func (i *Jfile) GetDestinationFile(FilesVersionName string) string {
	return i.GetTempleteName(i.DestinationFolder, FilesVersionName)
}

/*
 */
func (i *Jfile) GetDestinationFileName(OriginFolderName string, FilesVersionName string) string {
	finaldir := i.DestinationDir + "/" + OriginFolderName + "/" + FilesVersionName
	i.Log.Debug("DestinationDir", finaldir)
	if !i.FileExist(finaldir) {
		i.Log.Warn("file no exist", finaldir)
	}
	return finaldir
}

/*
 */
func (i *Jfile) GetTempleteName(OriginFolderName string, FilesVersionName string) string {
	finaldir := i.OrigHomedir + "/" + OriginFolderName + "/" + FilesVersionName
	i.Log.Debug("GetTempleteName", finaldir)
	if !i.FileExist(finaldir) {
		i.Log.Warn("file no exist", finaldir)
	}
	return finaldir
}

/*
	Desactiva el debug
*/
func (i *Jfile) LogDebugOff() {
	if i.Log != nil {
		i.Log.Debug("Debug off")
		i.Log.DebugOff()
	}
}

/*
	Configura por defecto de log
*/
func (i *Jfile) PrepareDefaultLog() {
	i.Log = jlog.PrepareLog(true, true, true)
}

/*
valida si un archivo existe
*/
/*
 */
func (i *Jfile) FileExist(dir string) bool {
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		i.Log.Debug("file exist", dir)
		return true
	}
	i.Log.Debug("file not exist", dir)
	return false
}

/*
 Crea un Folder y sino puede para el proceso
*/
func (i *Jfile) CreateFolderOrDie(Dir string) {
	i.CreateFolder(Dir, true)
}

/*
Crea un folder y sino puede no importa
*/
func (i *Jfile) CreateFolderAndContinue(Dir string) {
	i.CreateFolder(Dir, false)
}

/*
	Crear folder con parametro "die"
	esta funcion se usa en CreateFolderAndContinue y CreateFolderOrDie
*/
func (i *Jfile) CreateFolder(Dir string, die bool) {
	i.Log.Debug("Creating folder...", Dir)
	var err error
	if i.FileExist(Dir) {
		err = errors.New("This folder exist:" + Dir)
	} else {
		err = os.Mkdir(Dir, 0755)
	}
	if err == nil {
		i.Log.Debug("Folder ready:", Dir)
	} else {
		i.Log.Error(err.Error(), Dir)
		if die {
			i.Log.IsFatal(err)
		}
	}

}

/* crea una ruta completa de Dir ejemplo
/ruta1/ruta2/ruta3
si no existe crea todo
*/
func (i *Jfile) CreateDirAll(Dir string) {
	if !i.FileExist(Dir) {
		i.Log.IsFatal(os.MkdirAll(Dir, 0755))
	} else {
		err := errors.New("this folder exist:" + Dir)
		i.Log.Error(err.Error(), Dir)
	}

}

/* Crea un archivo y le agrega data
 */
func (i *Jfile) CreateFile(DestinationDir string, Data string) {
	i.Log.Debug("Creating file...", DestinationDir)
	i.Log.IsFatal(ioutil.WriteFile(DestinationDir, []byte(Data), 0644))
}

/* lee un archivo y retorna un string de lo leido
 */
func (i *Jfile) ReadFile(TemplateName string) string {
	i.Log.Debug("Reading file...", TemplateName)
	data, err := ioutil.ReadFile(TemplateName)
	i.Log.IsFatal(err)
	return string(data)
}

/* Copia un archivo usando operaciones del sistema
 */
func (i *Jfile) Copy(srcFileDir string, DestFileDir string) {
	srcFile, err := os.Open(srcFileDir)
	i.Log.IsFatal(err)
	defer srcFile.Close()

	destFile, err := os.Create(DestFileDir) // creates if file doesn't exist
	i.Log.IsFatal(err)
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	i.Log.IsFatal(err)

	err = destFile.Sync()
	i.Log.IsFatal(err)
}

/* lee un archivo y luego lo copia a otro
 */
func (i *Jfile) ReadAndCopy(srcFileDir string, DestFileDir string) {
	b, err := ioutil.ReadFile(srcFileDir)
	i.Log.IsFatal(err)
	err = ioutil.WriteFile(DestFileDir, b, 0644)
	i.Log.IsFatal(err)
}

/* Remplaza terminos en un archivo con data, los terminos a remplazar estan en el MapForReplace
 */
func (i *Jfile) UpdateFile(TemplateName string, MapForReplace map[string]string) {
	data := i.ReplaceTextInFile(TemplateName, MapForReplace)
	i.CreateFile(TemplateName, data)
}

/* Remplaza terminos en un archivo con data, los terminos a remplazar estan en el i.Map
 */
func (i *Jfile) UpdateFileWithInternalData(DestinationFileName string) {
	data := i.ReplaceTextInFile(DestinationFileName, i.Map)
	i.CreateFile(DestinationFileName, data)
}

/* Crea un archivo nuevo partiendo de un una plantilla y un arreglo de opciones a remplazar
 */
func (i *Jfile) NewFileforTemplate(NewName string, TemplateName string, MapForReplace map[string]string) {
	data := i.ReplaceTextInFile(TemplateName, MapForReplace)
	i.CreateFile(NewName, data)
}

/* Crea un archivo nuevo partiendo de un una plantilla y un arreglo de opciones a remplazar en el i.Map
 */
func (i *Jfile) NewFileForTemplateWithInternalData(NewName string, TemplateName string) {
	data := i.ReplaceTextInFile(TemplateName, i.Map)
	i.CreateFile(NewName, data)
}

/*
 */
func (i *Jfile) CopyTemplate(OriginFolderName string, FilesVersionName string, destFileName string) {
	i.Copy(i.GetTempleteName(OriginFolderName, FilesVersionName), destFileName)
}

/* lee un file y remplaza info en el luego lo retorna (no remplaza el archivo los cambios lo pone en memoria)
 */
func (i *Jfile) ReplaceTextInFile(TemplateName string, MapForReplace map[string]string) string {
	input := i.ReadFile(TemplateName)
	for key, value := range MapForReplace {
		input = strings.Replace(input, key, value, -1)
	}
	return input
}

/*añade al final del archivo un string
 */
func (i *Jfile) AppEndToFile(DestFileDir string, data string) {
	b, err := os.OpenFile(DestFileDir, os.O_APPEND|os.O_WRONLY, 0600)
	i.Log.IsFatal(err)
	defer b.Close()
	_, err = b.WriteString(data)
	i.Log.IsFatal(err)
}

/*
agrega al final de un archivo varios string
*/
func (i *Jfile) AppEndArrayToFile(DestFileDir string, datas []string) {
	b, err := os.OpenFile(DestFileDir, os.O_APPEND|os.O_WRONLY, 0600)
	i.Log.IsFatal(err)
	defer b.Close()
	for _, data := range datas {
		_, err = b.WriteString(data)
		i.Log.IsFatal(err)
	}
}

/*
usa el i.Map y va agregando data en el
*/
func (i *Jfile) AddMap(index string, value string) {
	if i.Map == nil {
		i.Map = make(map[string]string)
	}
	i.Map[index] = value
}

/*
agrega un valor a un archivo remplazando lo que consiga en "index" por el valor de "value"
*/
func (i *Jfile) AddValueToFile(index string, value string, Destinationfile string) {
	MapForReplace := make(map[string]string)
	MapForReplace[index] = value
	i.UpdateFile(Destinationfile, MapForReplace)
}

/* lee un archivo y pone su contenido en una linea usando el separador
 */
func (i *Jfile) ReadFileAndPutInLine(filePath string, separator string) string {

	readFile, err := os.Open(filePath)
	i.Log.IsFatal(err)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines string
	fileScanner.Scan()
	fileLines = fileScanner.Text()
	for fileScanner.Scan() {
		fileLines = fileLines + separator + fileScanner.Text()
	}
	readFile.Close()
	return fileLines
}

/*
	Add content to file if not exist match string
*/
func (i *Jfile) AddContentIfNotExist(DestineName string, MatchString string, ContentToAdd string) bool {
	word := i.ReadFile(DestineName)
	existe, _ := regexp.MatchString(MatchString, word)
	if !existe {
		i.AppEndToFile(DestineName, ContentToAdd)
	}
	return existe
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
func (i *Jfile) GetJsonFileWithStruct(jsonFileName string, WithStruct interface{}) {
	jsonFile, err := os.Open(jsonFileName)
	i.Log.IsFatal(err)
	defer jsonFile.Close()
	byteValue, err1 := ioutil.ReadAll(jsonFile)
	i.Log.IsFatal(err1)
	i.Log.IsFatal(json.Unmarshal([]byte(byteValue), WithStruct))
}
