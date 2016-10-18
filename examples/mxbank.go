package main

import (
	"fmt"
	"github.com/mxenabled/atrium-go/client"
	"github.com/mxenabled/atrium-go/models"
	"os"
	"time"
)

func getEnv(key string) string {
	value, isPresent := os.LookupEnv(key)

	if !isPresent {
		fmt.Println("You need to set the", key, "as an environment variable.")
		os.Exit(1)
	}

	return value
}

func main() {
	// Create a new Client
	client := &client.Client{
		ApiKey:   getEnv("API_KEY"),
		ClientId: getEnv("CLIENT_ID"),
		ApiURL:   "https://vestibule.mx.com",
	}

	// Create a User
	newUser := &models.User{
		Metadata: "info to store on the user",
	}
	user, err := client.CreateUser(newUser)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	fmt.Println("Created a user with guid:", user.Guid)

	// Get credentials for MX Bank
	credentials, err := client.ListCredentials("mxbank")
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	// Create login credentials
	loginCredentials := []*models.Credential{}
	for _, credential := range credentials {
		if credential.Type == "LOGIN" {
			username := &models.Credential{Guid: credential.Guid, Value: "atrium"}
			loginCredentials = append(loginCredentials, username)
		}

		if credential.Type == "PASSWORD" {
			password := &models.Credential{Guid: credential.Guid, Value: "anyPassword"}
			loginCredentials = append(loginCredentials, password)
		}
	}

	// Create a Member for the User with login credentials.
	// When we create a new Member, we'll attempt to aggregate data from MX Bank.
	newMember := &models.Member{
		Metadata:        "info to store on the member",
		InstitutionCode: "mxbank",
	}
	member, err := client.CreateMember(user.Guid, newMember, loginCredentials)
	if err != nil {
		fmt.Println("Error creating member:", err)
		return
	}

	maxAttempts := 10
	for {
		maxAttempts--

		if maxAttempts == 0 || member.Status == "HALTED" {
			fmt.Println(user.Guid, member.Guid, "Member agg failed!")
			return
		}

		if member.Status == "COMPLETED" {
			fmt.Println(user.Guid, member.Guid, "Member agg completed!")
			break
		}

		time.Sleep(time.Second)

		// Refresh the latest member status
		member, err = client.GetMember(user.Guid, member.Guid)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(user.Guid, member.Guid, "current status: ", member.Status)
	}

	// Get user Accounts after a commpleted aggregation.
	accounts, err := client.ListAccounts(user.Guid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.Guid, "Found", len(accounts), "accounts.")

	for _, account := range accounts {
		fmt.Println(user.Guid, account.Guid, account.Name)

		// Get transactions for an Account
		transactions, err := client.ListAccountTransations(user.Guid, account.Guid)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(user.Guid, "Found", len(transactions), "transactions.")

		for _, transaction := range transactions {
			fmt.Println(user.Guid, account.Guid, transaction.Guid, transaction.Description)
		}
	}

	// Delete the User
	err = client.DeleteUser(user.Guid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.Guid, "was deleted.")
}
