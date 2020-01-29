package executor

import (
	"log"
	"os/exec"
	"syscall"
)

func RunCmd(cmd []string, env map[string]string) int {
	command := exec.Command(cmd[0], cmd[1:]...)
	command.Env = prepareEnvVar(env)
	err := command.Start()
	if err != nil {
		return 1
	}
	if err = command.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return int(status.ExitCode)
			}
		} else {
			log.Fatal(err)
		}
	}
	return 0
}

func prepareEnvVar(envs map[string]string) []string {
	res := make([]string, len(envs))
	for key, val := range envs {
		res = append(res, key+"="+val)
	}
	return res
}
