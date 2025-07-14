package main

import (
	"Backend-POS/utils"
	"fmt"
	"log"
)

func main() {
	passwords := []string{"password", "admin123", "staff123"}

	fmt.Println("=== Password Hash Generator ===")
	for _, password := range passwords {
		hash, err := utils.HashPassword(password)
		if err != nil {
			log.Fatal("Error hashing password:", err)
		}

		fmt.Printf("Password: %s\n", password)
		fmt.Printf("Hash: %s\n", hash)
		fmt.Println("---")
	}

	testPassword := "password"
	hash, _ := utils.HashPassword(testPassword)
	isValid := utils.CheckPasswordHash(testPassword, hash)

	fmt.Printf("Testing - Password: %s, Valid: %t\n", testPassword, isValid)
}
