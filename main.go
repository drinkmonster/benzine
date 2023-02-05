package main

import (
	"fmt"
	"os"
	"os/exec"
)

func execute(cmd string, args string) {

	out, err := exec.Command(cmd, args).Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	// fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)
}

func list() {
	execute("echo", "Available packages:")
	execute("ls", "pkgs")
}

func firstTimeLaunch() {
	execute("mkdir", "~/.pkgs")
}

func help() {
	execute("echo", "hello")
}

func main() {
	arg := os.Args[1]
	arg2 := os.Args[2]
	if arg == "help" {
		help()
	}
	if arg == "install" && arg2 == "pfetch" {
		execute("bash", "pkgs/pfetch/template")
	}
}
