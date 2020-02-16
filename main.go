package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

var args = os.Args[1:]

func main() {
	if len(args) == 0 {
		log.Fatal("You need to provide a port number")
	}

	var port = args[0]

	if _, err := strconv.Atoi(port); err != nil {
		log.Fatalln("Could not parse port ", err)
	}

	linuxCommand := fmt.Sprintf("fuser -k %s/tcp", port)

	execute(exec.Command("bash", "-c", linuxCommand))
}


func execute(command *exec.Cmd) {
	var status syscall.WaitStatus
	var exitStatus int

	if err := command.Run(); err != nil {
		fmt.Println("Error while running the command", err.Error())

		if exitError, found := err.(*exec.ExitError); found {
			status = exitError.Sys().(syscall.WaitStatus)
			exitStatus = status.ExitStatus()
			log.Fatalf("Command exited with: %d", exitStatus)
		}

	} else {
		status = command.ProcessState.Sys().(syscall.WaitStatus)
		exitStatus = status.ExitStatus()
		log.Printf("Successfully killed process on port %s. Exit code: %d", args[0], exitStatus)
	}
}
