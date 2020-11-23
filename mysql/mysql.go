package mysql

import (
	"os/exec"

	"github.com/culturadevops/jgt/exe"
	"github.com/culturadevops/jgt/jio"
)

type Jmysql struct {
	Dns    string
	Root   string
	Pass   string
	Squema string
}

func SetVar(Dns string,Root string,Passs string, Squema string) Jmysql {
	return Jmysql{
		Dns:    "uatcupomania.c4rjguxjh6uc.us-east-1.rds.amazonaws.com",
		Root:   "root",
		Pass:   "xOGDOwsMsAN8mLTUSTLjic7I",
		Squema: "uatcupomania",
	}
}

func (j *Jmysql) Command(arg ...string) *exec.Cmd {
	arg = append([]string{"-u", j.Root, "-p" + j.Pass, "-h", j.Dns}, arg...)
	cmd := exec.Command("mysql", arg...)
	return cmd
}
func (j *Jmysql) Import(data string) {
	exe.RunWithData(j.Command(j.Squema), data)
}

func (j *Jmysql) Dump(folder string, squema string) {
	if squema != "" {
		j.Squema = squema
	}
	cmd := exec.Command("mysqldump", "--no-data", "-u", j.Root, "-p"+j.Pass, "--set-gtid-purged=OFF", "-h", j.Dns, j.Squema)
	jio.CreateFile(folder+"/"+j.Squema+".sql", exe.Run(cmd, false))

}

func (j *Jmysql) dumpSameHostMultiSquema(dir string, squemas []string) {
	for _, squema := range squemas {
		j.Dump(dir, squema)
	}
}
