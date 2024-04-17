package shared

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetScanner() int {
	var choice int
	fmt.Scanln(&choice)
	return choice
}

func GetScannerString() string {
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	return choice
}
