package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if _, err := os.Stat(".git"); err != nil && os.IsNotExist(err) {
		fmt.Println("fatal: not a git repository")
		return
	}

	cmd := exec.Command("git", "remote", "-v")
	var bs []byte
	bs, err := cmd.Output()
	if err != nil {
		fmt.Println("fatal: command output failed")
		return
	}

	if len(bs) == 0 {
		fmt.Println("fatal: not a remote repository")
		return
	}

	ss := strings.Split(string(bs), "\n")
	i1 := strings.Index(ss[0], "http")
	i2 := strings.Index(ss[0], ".git")
	url := ss[0][i1 : i2+4]
	cmd = exec.Command("open", url)
	cmd.Run()
}
