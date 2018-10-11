package main

import (
	"fmt"
	"github.com/mxenabled/atrium-go/client"
	"github.com/mxenabled/atrium-go/models"
	"os"
)

func getEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		fmt.Println("You need to set the", key, "as an environment variable.")
		os.Exit(1)
	}
	return value
}

func main() {
	client := &client.Client{
		ApiKey:   getEnv("API_KEY"),
		ClientId: getEnv("CLIENT_ID"),
		ApiURL:   "https://vestibule.mx.com",
	}

	newUser := &models.User{
		Metadata: "info to store on the user",
	}

	user, err := client.CreateUser(newUser)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	fmt.Println("Created a user with guid:", user.Guid)

	err = client.DeleteUser(user.Guid)
	if err != nil {
		fmt.Println("Error deleting user:", err)
		return
	}
	fmt.Println("User", user.Guid, "was deleted")
}
