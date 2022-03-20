package jfile

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
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

/* Agrega datos para usar las variables de folder home y origen
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

/* genera una ruta de una plantilla usando origin home y origin folder  de Jfile
 */
func (i *Jfile) GetOriginFile(OriginFileName string) string {
	return i.GetTempleteName(i.OrigFolder, OriginFileName)
}

/* genera ruta de destino usando destinohome y destion folder de el objeto jfile
 */
func (i *Jfile) GetDestinationFile(filesVersionName string) string {
	return i.GetTempleteName(i.DestinationFolder, filesVersionName)
}

/*
 */
func (i *Jfile) GetDestinationFileName(origFolderName string, filesVersionName string) string {
	finaldir := i.DestinationDir + "/" + origFolderName + "/" + filesVersionName
	i.Log.Debug("DestinationDir", finaldir)
	if !i.FileExist(finaldir) {
		i.Log.Warn("file no exist", finaldir)
	}
	return finaldir
}

/*
 */
func (i *Jfile) GetTempleteName(origFolderName string, filesVersionName string) string {
	finaldir := i.OrigHomedir + "/" + origFolderName + "/" + filesVersionName
	i.Log.Debug("GetTempleteName", finaldir)
	if !i.FileExist(finaldir) {
		i.Log.Warn("file no exist", finaldir)
	}
	return finaldir
}

/*
 */
func (i *Jfile) DebufOff() {
	if i.Log != nil {
		i.Log.IsDebug = false
	}
}

/*
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
 */
func (i *Jfile) CreateFolderOrDie(directorio string) {
	i.CreateFolder(directorio, true)
}

/*
 */
func (i *Jfile) CreateFolderAndContinue(directorio string) {
	i.CreateFolder(directorio, false)
}

/*
 */
func (i *Jfile) CreateFolder(directorio string, die bool) {
	i.Log.Debug("Inicio Creando archivo", directorio)
	if !i.FileExist(directorio) {
		err := os.Mkdir(directorio, 0755)
		i.Log.Error(err.Error(), directorio)
		if die {
			i.Log.IsFatal(err)
		}
	}
}

/* crea una ruta completa de directorio ejemplo
/ruta1/ruta2/ruta3
si no existe crea todo
*/
func (i *Jfile) CreateDirAll(directorio string) {
	if !i.FileExist(directorio) {
		i.Log.IsFatal(os.MkdirAll(directorio, 0755))
	}
}

/* Crea un archivo y le agrega data
 */
func (i *Jfile) CreateFile(rutaDestino string, data string) {
	i.Log.Debug("Inicio Creando archivo", rutaDestino)
	i.Log.IsFatal(ioutil.WriteFile(rutaDestino, []byte(data), 0644))
}

/* lee un archivo y retorna un string de lo leido
 */
func (i *Jfile) ReadFile(templateName string) string {
	i.Log.Debug("Inicio lectura de archivo", templateName)
	data, err := ioutil.ReadFile(templateName)
	i.Log.IsFatal(err)
	return string(data)
}

/* Copia un archivo usando operaciones del sistema
 */
func (i *Jfile) Copy(srcFileDir string, destFileDir string) {
	srcFile, err := os.Open(srcFileDir)
	i.Log.IsFatal(err)
	defer srcFile.Close()

	destFile, err := os.Create(destFileDir) // creates if file doesn't exist
	i.Log.IsFatal(err)
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	i.Log.IsFatal(err)

	err = destFile.Sync()
	i.Log.IsFatal(err)
}

/* lee un archivo y luego lo copia a otro
 */
func (i *Jfile) ReadAndCopy(srcFileDir string, destFileDir string) {
	b, err := ioutil.ReadFile(srcFileDir)
	i.Log.IsFatal(err)
	err = ioutil.WriteFile(destFileDir, b, 0644)
	i.Log.IsFatal(err)
}

/* modifica un archivo buscando algo
 */
func (i *Jfile) UpdateFile(templateName string, MapForReplace map[string]string) {
	data := i.ReplaceTextInFile(templateName, MapForReplace)
	i.CreateFile(templateName, data)
}

/*
 */
func (i *Jfile) UpdateFileWithInternalData(DestinationFileName string) {
	data := i.ReplaceTextInFile(DestinationFileName, i.Map)
	i.CreateFile(DestinationFileName, data)
}

/* Crea un archivo nuevo partiendo de un una plantilla y un arreglo de opciones a remplazar
 */
func (i *Jfile) NewFileforTemplate(newName string, templateName string, MapForReplace map[string]string) {
	data := i.ReplaceTextInFile(templateName, MapForReplace)
	i.CreateFile(newName, data)
}

/*
 */
func (i *Jfile) CopyTemplate(origFolderName string, filesVersionName string, destFileName string) {
	i.Copy(i.GetTempleteName(origFolderName, filesVersionName), destFileName)
}

/* remplaza info en un archivo luego lo pasa a una variable
 */
func (i *Jfile) ReplaceTextInFile(templateName string, MapForReplace map[string]string) string {
	input := i.ReadFile(templateName)
	for key, value := range MapForReplace {
		input = strings.Replace(input, key, value, -1)
	}
	return input
}

/*añade al final del archivo un string
 */
func (i *Jfile) AppEndToFile(destFileDir string, data string) {
	b, err := os.OpenFile(destFileDir, os.O_APPEND|os.O_WRONLY, 0600)
	i.Log.IsFatal(err)
	defer b.Close()
	_, err = b.WriteString(data)
	i.Log.IsFatal(err)
}

/*
 */
func (i *Jfile) AppEndArrayToFile(destFileDir string, datas []string) {
	b, err := os.OpenFile(destFileDir, os.O_APPEND|os.O_WRONLY, 0600)
	i.Log.IsFatal(err)
	defer b.Close()
	for _, data := range datas {
		_, err = b.WriteString(data)
		i.Log.IsFatal(err)
	}
}

/*
 */
func (i *Jfile) AddMap(index string, value string) {
	if i.Map == nil {
		i.Map = make(map[string]string)
	}
	i.Map[index] = value
}

/*
 */
func (i *Jfile) AddValueToFile(index string, value string, Destinationfile string) {
	MapForReplace := make(map[string]string)
	MapForReplace[index] = value
	i.UpdateFile(Destinationfile, MapForReplace)
}

/*
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
