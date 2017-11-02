package util

import (
	"blog/app/config"
	"os/exec"
	"fmt"
)

func Killprocess() {
	shell := "lsof -t -i tcp:" + config.Port
	out := execCommandLine(shell)

	res := "kill " + string(out)
	fmt.Println(res)

}

func execCommandLine(shell string) []byte {
	cmd := exec.Command("sh", "-c", shell)
	out, err := cmd.Output()
	CheckErr(err)
	return out
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
