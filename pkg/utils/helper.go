package utils

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func ExecuteCommand(name string, subname string, args ...string) (string, error) {
	args = append([]string{subname}, args...)

	cmd := exec.Command(name, args...)
	bytes, err := cmd.CombinedOutput()
	fmt.Println(err)
	return string(bytes), err
}

func Error(cmd *cobra.Command, args[] string, err error) {
	fmt.Fprint(os.Stderr,"exec %s args:%v error:%v\n", cmd.Name(), args, err)
	os.Exit(1)
}