
// Package main provides a Windows-WSL bridge for Docker commands
// allowing seamless execution of Docker commands from Windows through WSL.
package main

import (
	"fmt"
	"os"
	"os/exec"
)

// runDockerComposeCommand executes Docker Compose commands through WSL
// It forwards all command-line arguments to the Docker Compose CLI running in WSL
// and handles input/output streams appropriately.
//
// The function will:
// - Check if WSL is available on the system
// - Use WSL_DISTRO environment variable or default to Ubuntu
// - Forward all command line arguments to the Docker Compose command in WSL
// - Handle standard input/output/error streams
// - Exit with appropriate error codes if command fails
func runDockerComposeCommand() {
	// Get all command-line arguments
	args := os.Args[1:]

	// Check if wsl is available
	if _, err := exec.LookPath("wsl"); err != nil {
		fmt.Fprintln(os.Stderr, "Error: wsl is not installed or not in PATH")
		os.Exit(1)
	}

	// Use environment variable for WSL distribution or default to "Ubuntu"
	wslDistro := os.Getenv("WSL_DISTRO")
	if wslDistro == "" {
		wslDistro = "Ubuntu"
	}

	// Construct the WSL command
	wslCmd := exec.Command("wsl", append([]string{"-d", wslDistro, "docker-compose"}, args...)...)

	// Set up input/output streams
	wslCmd.Stdin = os.Stdin
	wslCmd.Stdout = os.Stdout
	wslCmd.Stderr = os.Stderr

	// Run the command
	err := wslCmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}