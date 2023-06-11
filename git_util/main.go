/*
Copyright © 2023 HAIDER ALI <haider.lee23@gmail.com>
*/
package git_util

import (
	"fmt"
	"os"
	"os/exec"
)

func Diff() string {
	app := "git"
	arg0 := "diff"
	arg1 := "--cached"

	cmd := exec.Command(app, arg0, arg1)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println("Git diff failed: " + err.Error())
		os.Exit(1)
	}

	return string(stdout)
}

func Commit(message string) {
	app := "git"
	arg0 := "commit"
	arg1 := "-m"

	cmd := exec.Command(app, arg0, arg1, message)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println("Git commit failed: " + err.Error())
		os.Exit(1)
	}

	fmt.Println(string(stdout))
}
