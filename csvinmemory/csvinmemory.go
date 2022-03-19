package csvinmemory

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/culturadevops/jgt/jlog"
)

type CsvInMemory struct {
	InMemory []map[string]string
	Log      *jlog.Jlog
	Head     []string
}

func (i *CsvInMemory) DebufOff() {
	if i.Log != nil {
		i.Log.IsDebug = false
	}
}
func (i *CsvInMemory) PrepareLog(IsDebug bool, PrinterLogs bool, PrinterScreen bool) {
	i.Log = &jlog.Jlog{
		IsDebug:       IsDebug,
		PrinterLogs:   PrinterLogs,
		PrinterScreen: PrinterScreen,
	}
	i.Log.SetInitProperty()
}
func (i *CsvInMemory) PrepareDefaultLog() {
	i.PrepareLog(true, true, true)

}
func (i *CsvInMemory) OpenFile(file string) *csv.Reader {
	f, err := os.Open(file)
	i.Log.IsFatal(err)
	r := csv.NewReader(f)
	return r
}
func (i *CsvInMemory) ReadHead(csvreader *csv.Reader) {
	var head []string
	i.Log.Debug("Inciando lectura de head", nil)
	record, err := csvreader.Read()
	if err == io.EOF {
		i.Log.Error(err.Error(), nil)
		return
	}
	for value := range record {
		head = append(head, record[value])
		i.Log.Debug("valor guardado", record[value])
	}
	i.Head = head
	i.Log.Debug("Terminando lectura de head", nil)
}
func (i *CsvInMemory) CreateCsvInMemory(head []string, r *csv.Reader) {
	var file map[string]string
	file = make(map[string]string)
	i.Log.Debug("Inciando lectura de CUERPO", nil)
	for {
		file = make(map[string]string)
		record, err := r.Read()
		i.Log.Debug("leyendo linea", record)
		if err == io.EOF {
			i.Log.Warn(err.Error(), nil)
			break
		}

		i.Log.IsFatal(err)

		h := 0
		for value := range record {
			if head[h] != "-" {
				file[head[h]] = record[value]
				//i.Log.Debug(head[h]+"->"+record[value], record[value])
			}
			h = h + 1
		}
		i.Log.Debug("Agregando a arreglo base la siguiente fila")
		i.Log.Debug("fila", file)

		i.InMemory = append(i.InMemory, file)
		i.Log.Debug("Terminando lectura de CUERPO", nil)
	}
}
