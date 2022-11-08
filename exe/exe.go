package exe

import (
	"bytes"
	"errors"
	"os/exec"
	"path/filepath"

	"github.com/culturadevops/jgt/jlog"
)

type Jexe struct {
	Log           *jlog.Jlog
	Arg           []string
	Cmd           *exec.Cmd
	Executable    string
	IsDebug       bool
	PrinterLogs   bool
	PrinterScreen bool
	ShowStd       bool
	ShowErr       bool
}

/*
	Desactiva el debug
*/
func (i *Jexe) LogDebugOff() {
	i.IsDebug = false
	if i.Log != nil {
		i.Log.DebugOff()
		i.Log.Debug("Debug off")
	}
}

/*
	Configura por defecto de log
*/
func (i *Jexe) PrepareDefaultLog() {
	i.IsDebug = true
	i.PrinterLogs = true
	i.PrinterScreen = true
	i.Log = jlog.PrepareLog(i.IsDebug, i.PrinterLogs, i.PrinterScreen)
}
func (i *Jexe) ConfigureLog(IsDebug bool, PrinterLogs bool, PrinterScreen bool) {
	i.IsDebug = IsDebug
	i.PrinterLogs = PrinterLogs
	i.PrinterScreen = PrinterScreen
	i.Log = jlog.PrepareLog(i.IsDebug, i.PrinterLogs, i.PrinterScreen)
}
func (i *Jexe) PrepareLog() {
	i.Log = jlog.PrepareLog(i.IsDebug, i.PrinterLogs, i.PrinterScreen)
}
func (i *Jexe) PrepareDefaultjExe(Executable string) {
	i.Executable = Executable
	i.ShowStd = true
	i.ShowErr = true
	i.IsDebug = true
	i.PrinterLogs = true
	i.PrinterScreen = true
	i.PrepareDefaultLog()

}
func (i *Jexe) PreparejExe(Executable string, ShowStd bool, ShowErr bool, IsDebug bool, PrinterLogs bool, PrinterScreen bool) {
	i.Executable = Executable
	i.ShowStd = ShowStd
	i.ShowErr = ShowErr
	i.IsDebug = IsDebug
	i.PrinterLogs = PrinterLogs
	i.PrinterScreen = PrinterScreen
	i.PrepareLog()
}
func (i *Jexe) CommandAndRun(withArgument bool, die bool) {
	i.Command(i.Executable, withArgument)
	i.Run(die)
}
func (i *Jexe) CommandInternal(withArgument bool) {
	i.Command(i.Executable, withArgument)
}

func (i *Jexe) Command(exectuble string, withArgument bool) {
	if withArgument {
		i.Cmd = exec.Command(exectuble, i.Arg...)
	} else {
		i.Cmd = exec.Command(exectuble)
	}
	i.Log.Debug("Commando:\n%s\n", i.Cmd)
}
func (i *Jexe) Run(die bool) (string, string, error) {
	var stdout, stderr bytes.Buffer
	i.Cmd.Stdout = &stdout
	i.Cmd.Stderr = &stderr
	err := i.Cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	if i.ShowStd {
		if outStr != "" {
			i.Log.Debug("exe output: \n%s\n", outStr)
		}
	}
	if errStr != "" {
		i.Log.IsErrorAndDie(errors.New(errStr), die)
	}
	if err != nil {
		i.Log.IsErrorAndDie(err, die)
	}
	return outStr, errStr, err
}
func (i *Jexe) GenerateAbsolutePath(FileName string) string {
	absPath, _ := filepath.Abs("./")
	return absPath + FileName
}
func (i *Jexe) AddParameterWithAbsolutePath(Paramterindex string, FileName string) []string {
	FileName = i.GenerateAbsolutePath(FileName)
	i.AddParameter(Paramterindex, FileName)
	return i.Arg
}
func (i *Jexe) AddParameter(Index string, Value string) []string {
	i.Arg = append([]string{Index, Value}, i.Arg...)
	return i.Arg
}
func (i *Jexe) Addflag(flag string) []string {
	i.Arg = append([]string{flag}, i.Arg...)
	return i.Arg
}
func (i *Jexe) RunWithData(data string, flagStdOut bool, die bool) (string, string, error) {
	buffer := bytes.Buffer{}
	buffer.Write([]byte(data))
	i.Cmd.Stdin = &buffer
	return i.Run(die)
}
