package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var output1, _ = exec.Command("is").Output()
	fmt.Printf("-> is \n%s\n", string(output1))

	var output2, _ = exec.Command("pwd").Output()
	fmt.Printf("-> pwd\n%s\n", string(output2))

	var output3, _ = exec.Command("git", "config", "user.name").Output()
	fmt.Printf(" -> git config user.name\n%s\n", string(output3))
}
