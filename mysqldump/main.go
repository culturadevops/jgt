package jmysqldump

import (
	"github.com/culturadevops/jgt/exe"
)

type Jzip struct {
	Jexe     *exe.Jexe
	Dns      string
	Root     string
	Pass     string
	Table    string
	DataBase string
}

func (j *Jzip) PrepareInit() {
	j.Jexe = new(exe.Jexe)
	j.Jexe.PrepareDefaultjExe("mysqldump")
}
func (j *Jzip) PrepareInSilence() {
	j.Jexe = new(exe.Jexe)
	j.Jexe.PrepareDefaultWithLogSilence("mysqldump")
}
func (j *Jzip) Dump(Die bool, DataBase string, Table string, filename string) {
	j.Jexe.Die = Die
	j.Jexe.ExecuteWithArg("-u", j.Root, "-p"+j.Pass, "-h", j.Dns, DataBase, Table, ">", filename)
}
