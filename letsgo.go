package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("./letsgo.sh")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
