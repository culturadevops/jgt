package exe

import (
	"bytes"
	"fmt"
	"os/exec"
)

func Run(cmd *exec.Cmd, flagStdOut bool) string {
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	if flagStdOut {
		fmt.Printf("out:\n%s\n", outStr)
	}
	fmt.Printf("err:\n%s\n", errStr)
	return outStr
}

func RunWithData(cmd *exec.Cmd, data string) {
	buffer := bytes.Buffer{}
	buffer.Write([]byte(data))
	cmd.Stdin = &buffer
	Run(cmd, false)
}
