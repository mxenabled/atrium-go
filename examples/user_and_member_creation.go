package main

import (
	"fmt"
	"github.com/mxenabled/atrium-go/client"
	"github.com/mxenabled/atrium-go/models"
)

func main() {
	// Create a new Client
	client := &client.Client{
		ApiKey:   "YOUR_MX_API_KEY",
		ClientId: "YOUR_MX_CLIENT_ID",
		ApiURL:   "https://vestibule.mx.com",
	}

	fmt.Println("\n* Creating test user *")
	newUser := &models.User{
		Metadata:   "some metadata",
		IsDisabled: false,
	}
	user, err := client.CreateUser(newUser)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	fmt.Println("Created User:", user.Guid)

	fmt.Println("\n* Listing institutions with query string \"bank\" *")
	institutions, listError := client.ListInstitutions("bank")
	if listError != nil {
		fmt.Println("Error listing institutions:", listError)
		return
	}
	for _, institution := range institutions {
		fmt.Println("Name:", institution.Name, ": institution code =", institution.Code)
	}

	fmt.Println("\n* Reading institution \"mxbank\" *")
	institution, readError := client.GetInstitution("mxbank")
	if err != nil {
		fmt.Println("Error reading institution:", readError)
		return
	}
	fmt.Println("Read Institution:", institution.Name)

	fmt.Println("\n* Creating member *")
	loginCredentials := []*models.Credential{}
	username := &models.Credential{Guid: "CRD-9f61fb4c-912c-bd1e-b175-ccc7f0275cc1", Value: "test_atrium"}
	loginCredentials = append(loginCredentials, username)
	password := &models.Credential{Guid: "CRD-e3d7ea81-aac7-05e9-fbdd-4b493c6e474d", Value: "password"}
	loginCredentials = append(loginCredentials, password)

	newMember := &models.Member{
		InstitutionCode: "mxbank",
	}
	member, memberError := client.CreateMember(user.Guid, newMember, loginCredentials)
	if err != nil {
		fmt.Println("Error creating member:", memberError)
		return
	}
	fmt.Println("Created member: ", member.Guid)

	fmt.Println("\n* Deleting test user *")
	deleteError := client.DeleteUser(user.Guid)
	if err != nil {
		fmt.Println("Error deleting user:", deleteError)
		return
	}
	fmt.Println("Deleted user:", user.Guid)
}
