package executor

import (
	"log"
	"os/exec"
	"syscall"
)

func RunCmd(cmd []string, env map[string]string) int {
	command := exec.Command(cmd[0], cmd[1:len(cmd)-1]...)
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
			log.Fatal("")
		}
	}
	return 0
}
