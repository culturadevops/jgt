package jlog

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Jlog struct {
	IsDebug       bool
	PrinterLogs   bool
	PrinterScreen bool
	LogInfo       *log.Logger
	LogError      *log.Logger
	DirLogs       string //os.Mkdir("logs", 0755)
	DirErroLogs   string //os.Mkdir("logs/error", 0755)
	LogFileName   string
}

func splitLast(file string) string {
	spliting := strings.Split(file, "/")
	x := len(spliting)
	return spliting[x-1]
}
func (i *Jlog) Write(typems string, format string, a ...interface{}) string {
	var ok bool
	ok = true
	var file, infofile string
	var line, actualline int
	var debuging, text string
	for in := 0; ok == true; in++ {
		_, file, actualline, ok = runtime.Caller(in)
		proctemp := splitLast(file)
		if proctemp == "proc.go" {
			break
		}
		infofile = proctemp
		debuging = debuging + " (" + proctemp + ":" + strconv.Itoa(actualline) + ")"

		line = actualline
		if i.IsDebug == true {
			infofile = debuging
		}
	}

	info := fmt.Sprintf(format+"-", a...)
	if i.IsDebug == true {
		text = fmt.Sprintf("%s: %v", infofile, info)
	} else {
		text = fmt.Sprintf("%s:%d: %v", infofile, line, info)
	}

	//fmt.Println(typems, text)
	return typems + text

}
func (i *Jlog) Debug(format string, a ...interface{}) {
	if i.IsDebug {
		texto := i.Write("[DEBUG]:", format, a...)
		color.White(texto)
		if i.PrinterLogs == true {
			i.LogInfo.Println(texto)
		}
	}
}
func (i *Jlog) Fatal(format string, a ...interface{}) {
	i.Error(format, a)
	os.Exit(1)
}
func (i *Jlog) IsFatal(err error) {
	if err != nil {
		i.Error(err.Error(), nil)
		os.Exit(1)
	}

}

func (i *Jlog) Error(format string, a ...interface{}) {
	texto := i.Write("[Error]:", format, a...)
	if i.PrinterScreen == true {
		color.Red(texto)
	}
	if i.PrinterLogs == true {
		i.LogError.Println(texto)
		i.LogInfo.Println(texto)
	}
}
func (i *Jlog) Info(format string, a ...interface{}) {
	texto := i.Write("[Info]:", format, a...)
	if i.PrinterScreen == true {
		color.White(texto)
	}
	if i.PrinterLogs == true {
		i.LogInfo.Println(texto)
	}

}
func (i *Jlog) Warn(format string, a ...interface{}) {
	texto := i.Write("[Warn]:", format, a...)
	if i.PrinterScreen == true {
		color.Yellow(texto)
	}
	if i.PrinterLogs == true {
		i.LogInfo.Println(texto)
	}
}

func (i *Jlog) SetInitProperty() {
	if i.PrinterLogs == true {
		if i.DirLogs == "" {
			i.DirLogs = "logs"
		}
		if i.DirErroLogs == "" {
			i.DirErroLogs = "logs/error"
		}
		if i.LogFileName == "" {
			i.LogFileName = "run"
		}
		os.Mkdir(i.DirLogs, 0755)
		os.Mkdir(i.DirErroLogs, 0755)
		logFile, err := os.OpenFile("./"+i.DirLogs+"/"+i.LogFileName+"_"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
		if err != nil {
			log.Fatalln("open log file failed", err)
		}
		logFileError, err1 := os.OpenFile("./"+i.DirErroLogs+"/"+i.LogFileName+"_"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
		if err1 != nil {
			log.Fatalln("open log file 'error' failed ", err)
		}
		//日志
		i.LogInfo = log.New(io.MultiWriter(logFile), "", log.Ldate|log.Ltime)       //LogInfo.Println(1, 2, 3)
		i.LogError = log.New(io.MultiWriter(logFileError), "", log.Ldate|log.Ltime) //LogError.Println(4, 5, 6)
	}
}
