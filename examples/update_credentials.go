package main

import (
	"fmt"
	"github.com/mxenabled/atrium-go/client"
	"github.com/mxenabled/atrium-go/models"
	"os"
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

	fmt.Println("\n* Creating user and member with \"DENIED\" aggregation status *")
	newUser := &models.User{
		Metadata:   "some metadata",
		IsDisabled: false,
	}
	user, err := client.CreateUser(newUser)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	fmt.Println("Created user:", user.Guid)

	loginCredentials := []*models.Credential{}
	username := &models.Credential{Guid: "CRD-9f61fb4c-912c-bd1e-b175-ccc7f0275cc1", Value: "test_atrium"}
	loginCredentials = append(loginCredentials, username)
	password := &models.Credential{Guid: "CRD-e3d7ea81-aac7-05e9-fbdd-4b493c6e474d", Value: "INVALID"}
	loginCredentials = append(loginCredentials, password)

	newMember := &models.Member{
		InstitutionCode: "mxbank",
	}
	member, memberError := client.CreateMember(user.Guid, newMember, loginCredentials)
	if err != nil {
		fmt.Println("Error creating member:", memberError)
		return
	}
	fmt.Println("Created member:", member.Guid)
	member = member
	time.Sleep(time.Second)

	fmt.Println("\n* Retrieving member aggregation status *")
	memberStatus, readError := client.GetMemberStatus(user.Guid, member.Guid)
	if err != nil {
		fmt.Println("Error getting member status:", readError)
		return
	}
	fmt.Println("Member aggregation status:", memberStatus.Status)

	fmt.Println("\n* Updating credentials *")
	credentials, listError := client.ListCredentials("mxbank")
	if err != nil {
		fmt.Println("Error creating user:", listError)
		return
	}

	updatedCredentials := []*models.Credential{}
	userName := &models.Credential{Guid: credentials[0].Guid, Value: "test_atrium"}
	updatedCredentials = append(updatedCredentials, userName)
	passWord := &models.Credential{Guid: credentials[1].Guid, Value: "password"}
	updatedCredentials = append(updatedCredentials, passWord)

	updatedMember := &models.Member{
		Guid: member.Guid,
	}

	_, updateError := client.UpdateMember(user.Guid, updatedMember, updatedCredentials)
	if err != nil {
		fmt.Println("Error updating member:", updateError)
		return
	}
	time.Sleep(time.Second)

	fmt.Println("\n* Retrieving member aggregation status *")
	memStatus, readErr := client.GetMemberStatus(user.Guid, member.Guid)
	if err != nil {
		fmt.Println("Error getting member status:", readErr)
		return
	}
	fmt.Println("Member aggregation status:", memStatus.Status)

	fmt.Println("\n* Deleting test user *")
	deleteError := client.DeleteUser(user.Guid)
	if err != nil {
		fmt.Println("Error deleting user:", deleteError)
		return
	}
	fmt.Println("Deleted user:", user.Guid)
}
