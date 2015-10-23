package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"os/exec"
)

func main() {
	out, err := exec.Command("git", "branch", "-a").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The branches are:\n%s\n", out)
}
