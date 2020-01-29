package executor

import (
	"os"
	"os/exec"
)

func RunCmd(cmd []string, env map[string]string) int {
	command := exec.Command(cmd[0], cmd[1:]...)
	command.Env = prepareEnvVar(env)
	command.Stdout = os.Stdout
	command.Stdin = os.Stdin

	if err := command.Run(); err != nil {
	}

	return 0
}

func prepareEnvVar(envs map[string]string) []string {
	res := make([]string, 0, len(envs))
	for key, val := range envs {
		res = append(res, key+"="+val)
	}
	return res
}
