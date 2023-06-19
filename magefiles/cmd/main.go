package cmd

import (
	"fmt"
	"os/exec"

	"github.com/lukeshay/pms/magefiles/sysexit"
)

func Exec(args ...string) string {
	out, err := exec.Command(args[0], args[1:]...).Output()

	fmt.Println(string(out))

	if err != nil {
		panic(sysexit.Os(err))
	}

	return string(out)
}
