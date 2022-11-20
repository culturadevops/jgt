package jzip

import (
	"github.com/culturadevops/jgt/exe"
	"github.com/culturadevops/jgt/jlog"
)

type Jzip struct {
	Jexe *exe.Jexe
}

func (g *Jzip) PrepareInit() {
	g.Jexe = new(exe.Jexe)
	g.Jexe.PrepareDefaultjExe("zip")
}
func (i *Jzip) ConfigureInitLog(IsDebug bool, PrinterLogs bool, PrinterScreen bool) {
	i.Jexe.Log = jlog.PrepareLog(IsDebug, PrinterLogs, PrinterScreen)
}
func (g *Jzip) PrepareInSilence() {
	g.Jexe = new(exe.Jexe)
	g.Jexe.PrepareDefaultWithLogSilence("zip")
}
func (g *Jzip) ZipFolder(NameZip string, FolderZip string, Die bool) {
	g.Jexe.Die = Die
	g.Jexe.ExecuteWithArg("-r", NameZip, FolderZip)
}
