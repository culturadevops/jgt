package jio

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

/*
valida si un archivo existe
*/
func IsFileExist(dir string) bool {
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		return true
	}
	return false
}
func CreateDir(directorio string) {

	if _, err := os.Stat(directorio); os.IsNotExist(err) {
		err = os.Mkdir(directorio, 0755)
		//MkdirAll
		if err != nil {
			// Aquí puedes manejar mejor el error, es un ejemplo
			panic(err)
		}
	}
}

//revisar
func CreateDirAll(directorio string) {

	if _, err := os.Stat(directorio); os.IsNotExist(err) {
		err = os.MkdirAll(directorio, 0755)

		if err != nil {
			// Aquí puedes manejar mejor el error, es un ejemplo
			panic(err)
		}
	}
}

// Crea un archivo
func CreateFile(rutaDestino string, data string) {
	err := ioutil.WriteFile(rutaDestino, []byte(data), 0644)
	if err != nil {
		panic(err)
	}
}

func check(err error) {
	if err != nil {
		//	fmt.Println("Error : %s", err.Error())
		os.Exit(1)
	}
}

//	lee un archivo y retorna un string de lo leido
//"path/filepath" templateName, _ := filepath.Abs(templateName)
func ReadFile(templateName string) string {
	data, _ := ioutil.ReadFile(templateName)
	return string(data)
}

//Copia un archivo usando operaciones del sistema
func Copy(srcFileDir string, destFileDir string) {
	srcFile, err := os.Open(srcFileDir)
	check(err)
	defer srcFile.Close()

	destFile, err := os.Create(destFileDir) // creates if file doesn't exist
	check(err)
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile) // check first var for number of bytes copied
	check(err)

	err = destFile.Sync()
	check(err)
}

//lee un archivo y luego lo copia a otro
func ReadAndCopy(srcFileDir string, destFileDir string) {

	b, err := ioutil.ReadFile(srcFileDir)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(destFileDir, b, 0644)
	if err != nil {
		panic(err)
	}

}

// modifica un archivo buscando algo
func ChancarFile(templateName string, MapForReplace map[string]string) {
	data := ReplaceTextInFile(templateName, MapForReplace)
	CreateFile(templateName, data)
}

// Crea un archivo nuevo partiendo de un una plantilla y un arreglo de opciones a remplazar
func NewFileforTemplate(newName string, templateName string, MapForReplace map[string]string) {
	data := ReplaceTextInFile(templateName, MapForReplace)
	CreateFile(newName, data)
}

//remplaza info en un archivo luego lo pasa a una variable
func ReplaceTextInFile(templateName string, MapForReplace map[string]string) string {
	input := ReadFile(templateName)
	for key, value := range MapForReplace {
		input = strings.Replace(input, key, value, -1)
	}
	return input
}

//añade al final del archivo un string
func AppEndToFile(destFileDir string, data string) {
	b, err := os.OpenFile(destFileDir, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer b.Close()
	if _, err = b.WriteString(data); err != nil {
		panic(err)
	}
}

//falta probar....... agrega muchas lineas nuevas
func AppArrayEndToFile(destFileDir string, datas []string) {
	for _, data := range datas {
		AppEndToFile(destFileDir, data)
	}

}
