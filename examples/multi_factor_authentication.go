package main

import (
	"fmt"
	"github.com/mxenabled/atrium-go/client"
	"github.com/mxenabled/atrium-go/models"
	"time"
)

func main() {
	// Create a new Client
	client := &client.Client{
		ApiKey:   "YOUR_MX_API_KEY",
		ClientId: "YOUR_MX_CLIENT_ID",
		ApiURL:   "https://vestibule.mx.com",
	}

	fmt.Println("\n* Creating user and member with \"CHALLENGED\" aggregation status *")
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
	password := &models.Credential{Guid: "CRD-e3d7ea81-aac7-05e9-fbdd-4b493c6e474d", Value: "challenge"}
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

	fmt.Println("\n* MFA Challenge *")
	challenges, chalError := client.GetMemberChallenges(user.Guid, member.Guid)
	if err != nil {
		fmt.Println("Error getting member challenge", chalError)
		return
	}
	for _, challenge := range challenges {
		fmt.Println("Challenge Guid: ", challenge.Guid)
		fmt.Println("Challenge Label: ", challenge.Label)
	}

	challengeResponses := []*models.Challenge{}
	challengeOne := &models.Challenge{Guid: challenges[0].Guid, Value: "correct"}
	challengeResponses = append(challengeResponses, challengeOne)

	member, resumeError := client.ResumeMember(user.Guid, member.Guid, challengeResponses)
	if err != nil {
		fmt.Println("Error resuming member:", resumeError)
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
