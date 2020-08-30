package jenkinscli

import (
	"os/exec"
	"path/filepath"

	"github.com/culturadevops/jgt/exe"
)

type JbaseJenkins struct {
	JenkinscliDir string
	ServerURL     string
	Port          string
	User          string
	Pass          string
	FinalPath     string // donde se ejecutara el comando
	FlagStdOut    bool
}

//fmt.Printf("len=%d cap=%d %v\n", len(arg), cap(arg), arg)
func New(JenkinscliDir string, ServerURL string, Port string, User string, Pass string, FinalPath string) *JbaseJenkins {
	j := &JbaseJenkins{}
	j.JenkinscliDir = JenkinscliDir
	j.ServerURL = ServerURL
	j.Port = Port
	j.User = User
	j.Pass = Pass
	j.FinalPath = FinalPath
	j.FlagStdOut = false
	return j
}
func (j *JbaseJenkins) SetCredentials(User string, Pass string) {
	j.User = User
	j.Pass = Pass
}
func (j *JbaseJenkins) BuildUrl() string {
	return "http://" + j.User + ":" + j.Pass + "@" + j.ServerURL + ":" + j.Port
}

func (j *JbaseJenkins) Command(arg ...string) *exec.Cmd {
	arg = append([]string{"-jar", j.JenkinscliDir, "-s", j.BuildUrl()}, arg...)

	cmd := exec.Command("java", arg...)
	if j.FinalPath != "" {
		absPath, _ := filepath.Abs(j.FinalPath)
		cmd.Dir = absPath
	}
	return cmd
}
func (j *JbaseJenkins) jenkinsExec(arg ...string) {
	exe.Run(j.Command(arg...), j.FlagStdOut)
}
func (j *JbaseJenkins) CreateJob(jobName string, data string) {

	exe.RunWithData(j.Command("create-job", jobName), data)
}
func (j *JbaseJenkins) CreateJobOtherServer(url string, jobName string, data string) {
	j.ServerURL = url
	j.CreateJob(jobName, data)
}

func (j *JbaseJenkins) CreateSameJobOnMultiServer(urlServer []string, jobName string, data string) {
	for _, value := range urlServer {
		j.CreateJobOtherServer(value, jobName, data)
	}
}

func (j *JbaseJenkins) GetJob(jobName string) {
	j.jenkinsExec("get-job", jobName)
}

func (j *JbaseJenkins) BuildJob(jobName string, arg ...string) {
	arg = append([]string{"build", jobName, "-p"}, arg...)
	j.jenkinsExec(arg...)

}
