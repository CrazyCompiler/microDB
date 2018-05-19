package repl

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func handleMetaCommand(command string) {
	if strings.Contains(command,".exit") {
		os.Exit(0)
	}else{
		fmt.Printf("Unrecognized command %s \n", command);
	}
}

func main() {
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		command, _ := reader.ReadString('\n')
		handleMetaCommand(command);
	}
}
