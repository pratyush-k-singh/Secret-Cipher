package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Authenticate() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	return username == "pinnacle" && password == "secrecy"
}
