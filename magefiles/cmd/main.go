package cmd

import (
	"os/exec"

	"github.com/lukeshay/pms/magefiles/sysexit"
)

func Exec(args ...string) {
	out, err := exec.Command(args[0], args[1:]...).Output()

	print(string(out))

	if err != nil {
		panic(sysexit.Os(err))
	}
}
