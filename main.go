package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"cipher_interpreter/cipher"
    "cipher_interpreter/config"
    "cipher_interpreter/storage"
    "cipher_interpreter/utils"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Authenticate user
	if !utils.Authenticate() {
		fmt.Println("Authentication failed. Exiting.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your message: ")
	message, _ := reader.ReadString('\n')
	message = strings.TrimSpace(message)

	fmt.Print("Enter complexity level (1-5): ")
	var complexity int
	fmt.Scanln(&complexity)

	cipherText, key := cipher.GenerateCipher(message, complexity)
	storage.StoreCipherKey(key)

	fmt.Printf("Cipher Text: %s\n", cipherText)
	fmt.Println("Keys stored for 10 minutes.")

	go storage.ExpireCipherKeys(10 * time.Minute)

	for {
		fmt.Print("Decipher or Cipher? ")
		action, _ := reader.ReadString('\n')
		action = strings.TrimSpace(strings.ToLower(action))

		if action == "decipher" {
			fmt.Print("Enter cipher text: ")
			inputCipher, _ := reader.ReadString('\n')
			inputCipher = strings.TrimSpace(inputCipher)
			originalText := cipher.Decipher(inputCipher)
			fmt.Printf("Original Text: %s\n", originalText)
		} else if action == "cipher" {
			fmt.Print("Enter plain text: ")
			plainText, _ := reader.ReadString('\n')
			plainText = strings.TrimSpace(plainText)
			newCipher, newKey := cipher.GenerateCipher(plainText, complexity)
			storage.StoreCipherKey(newKey)
			fmt.Printf("Cipher Text: %s\n", newCipher)
		} else {
			fmt.Println("Invalid action.")
		}
	}
}
