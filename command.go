package common

import (
	"io/ioutil"
	"os/exec"
)

// CommandDefault 标准命令行
func CommandDefault(cmd string) (string, error) {
	c := exec.Command(cmd)
	stdout, err := c.StdoutPipe()
	if err != nil {
		return "", err
	}

	if err := c.Start(); err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}

	if err := c.Wait(); err != nil {
		return "", err
	}

	return string(data), nil
}

// CommandWithPipeline 带管道的命令行语句
func CommandWithPipeline(cmd string) (string, error) {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
