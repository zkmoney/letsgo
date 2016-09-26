package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// cmd := exec.Command("./letsgo.sh")
	// cmd.Stdin = os.Stdin
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// if err := cmd.Run(); err != nil {
	// 	log.Fatal(err)
	// }

	r := Runner{
		stdin:  os.Stdin,
		stdout: os.Stdout,
		buf:    []byte{},
	}

	var gopath, username string
	gopath = "testing/hi/sup/go"

	r.makeUserRepo(gopath + "/src/github.com/" + username)
}

type Runner struct {
	stdin  io.Reader
	stdout io.Writer
	buf    []byte
}

func (r Runner) makeUserRepo(dir string) {
	fmt.Fprintln(r.stdout, "What's your github handle?")
	r.read()
	if err := os.MkdirAll(dir, os.ModeDir|0755); err != nil {
		panic(err)
	}
	fmt.Fprintln(r.stdout, dir+"is all set\n")
}

func (r Runner) read() []byte {
	scanner := bufio.NewScanner(r.stdin)
	scanner.Scan()
	return scanner.Bytes()
}
