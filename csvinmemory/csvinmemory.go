package csvinmemory

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/culturadevops/jgt/jlog"
)

type CsvInMemory struct {
	InMemory []map[string]string
	log      *jlog.Jlog
	Head     []string
}

func (i *CsvInMemory) PrepareLog(IsDebug bool, PrinterLogs bool, PrinterScreen bool) {
	i.log = &jlog.Jlog{
		IsDebug:       IsDebug,
		PrinterLogs:   PrinterLogs,
		PrinterScreen: PrinterScreen,
	}
	i.log.SetInitProperty()
}
func (i *CsvInMemory) PrepareDefaultLog() {
	i.PrepareLog(true, true, true)

}
func (i *CsvInMemory) OpenFile(file string) *csv.Reader {
	f, err := os.Open(file)
	i.log.IsFatal(err)
	r := csv.NewReader(f)
	return r
}
func (i *CsvInMemory) ReadHead(csvreader *csv.Reader) {
	var head []string
	i.log.Debug("Inciando lectura de head", nil)
	record, err := csvreader.Read()
	if err == io.EOF {
		i.log.Error(err.Error(), nil)
		return
	}
	for value := range record {
		head = append(head, record[value])
		i.log.Debug("valor guardado", record[value])
	}
	i.Head = head
	i.log.Debug("Terminando lectura de head", nil)
}
func (i *CsvInMemory) CreateCsvInMemory(head []string, r *csv.Reader) {
	var file map[string]string
	file = make(map[string]string)
	i.log.Debug("Inciando lectura de CUERPO", nil)
	for {
		file = make(map[string]string)
		record, err := r.Read()
		i.log.Debug("leyendo linea", record)
		if err == io.EOF {
			i.log.Warn(err.Error(), nil)
			break
		}

		i.log.IsFatal(err)

		h := 0
		for value := range record {
			if head[h] != "-" {
				file[head[h]] = record[value]
				//i.log.Debug(head[h]+"->"+record[value], record[value])
			}
			h = h + 1
		}
		i.log.Debug("Agregando a arreglo base la siguiente fila")
		i.log.Debug("fila", file)

		i.InMemory = append(i.InMemory, file)
		i.log.Debug("Terminando lectura de CUERPO", nil)
	}
}
