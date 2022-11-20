package mysql

import (
	"github.com/culturadevops/jgt/exe"
)

type Jmysql struct {
	Dns    string
	Root   string
	Pass   string
	Squema string
	Jexe   *exe.Jexe
}

func (g *Jmysql) SetVar(Dns string, Root string, Pass string, Squema string) {

	g.Dns = Dns
	g.Root = Root
	g.Pass = Pass
	g.Squema = Squema
	g.PrepareInit()

}
func (g *Jmysql) PrepareInit() {
	g.Jexe = new(exe.Jexe)
	g.Jexe.PrepareDefaultjExe("mysql")
}
func (g *Jmysql) PrepareInSilence() {
	g.Jexe = new(exe.Jexe)
	g.Jexe.PrepareDefaultWithLogSilence("mysql")
}
func (j *Jmysql) CreateConection(arg ...string) []string {
	arg1 := append([]string{"-u", j.Root, "-p" + j.Pass, "-h", j.Dns}, arg...)
	return arg1
}

func (j *Jmysql) Command(arg ...string) {
	arg1 := append(j.CreateConection(), arg...)
	j.Jexe.ExecuteWithArg(arg1...)
}
func (j *Jmysql) Import(data string) {
	j.Jexe.ExecuteWithArgAndData(data, j.Squema)
}

/*
//"--set-gtid-purged=OFF"//"--no-data"
func (j *Jmysql) Dump(folder string, squema string, arg ...string) {
	if squema != "" {
		j.Squema = squema
	}
	arg = append([]string{"-u", j.Root, "-p" + j.Pass, "-h", j.Dns, j.Squema}, arg...)
	cmd := exec.Command("mysqldump", arg...)
	jio.CreateFile(folder+"/"+j.Squema+".sql", exe.Run(cmd, false))

}
func (j *Jmysql) DumpNoData(folder string, squema string) {
	j.Dump(folder, squema, "--no-data")
}
func (j *Jmysql) DumpSetGtidPurgedoff(folder string, squema string) {
	j.Dump(folder, squema, "--set-gtid-purged=OFF")
}
func (j *Jmysql) DumpSameHostMultiSquema(dir string, squemas []string) {
	for _, squema := range squemas {
		j.Dump(dir, squema)
	}
}*/
