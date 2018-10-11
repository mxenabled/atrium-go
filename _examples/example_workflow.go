package main

import (
	"bufio"
	"fmt"
	"github.com/mxenabled/atrium-go/client"
	"github.com/mxenabled/atrium-go/models"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var counter = 0

func getEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
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

	userGUID := ""
	memberGUID := ""
	endUserPresent := ""

	br := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter in user GUID. If not yet created just press enter key: ")
	userGUID, _ = br.ReadString('\n')
	userGUID = strings.TrimSuffix(userGUID, "\n")

	fmt.Println("\nPlease enter in member GUID. If not yet created just press enter key: ")
	memberGUID, _ = br.ReadString('\n')
	memberGUID = strings.TrimSuffix(memberGUID, "\n")

	fmt.Println("\nPlease enter in if end user is present (true or false): ")
	endUserPresent, _ = br.ReadString('\n')
	endUserPresent = strings.TrimSuffix(endUserPresent, "\n")

	if userGUID == "" && memberGUID != "" {
		fmt.Println("\nMust include user GUID when member GUID is entered.")
		os.Exit(0)
	}

	if userGUID == "" && endUserPresent == "true" {
		fmt.Println("\n* Creating user *")

		fmt.Println("\nPlease enter in an unique id: ")
		identifier, _ := br.ReadString('\n')
		identifier = strings.TrimSuffix(identifier, "\n")

		newUser := &models.User{
			Identifier: identifier,
			IsDisabled: false,
		}
		user, err := client.CreateUser(newUser)
		if err != nil {
			fmt.Println("Error creating user:", err)
			return
		}
		userGUID = user.Guid
		fmt.Println("\nCreated user:", userGUID)
	}

	if memberGUID != "" && endUserPresent == "true" {
		_, err := client.AggregateMember(userGUID, memberGUID)
		if err != nil {
			fmt.Println("Error aggregating member:", err)
			return
		}

		checkJobStatus(client, userGUID, memberGUID)
	} else if memberGUID != "" {
		readAggregationData(client, userGUID, memberGUID)
	} else if endUserPresent == "true" {
		fmt.Println("\n* Creating member *")

		fmt.Println("\n* Listing top 15 institutions *")
		institutions, err := client.ListInstitutions("")
		if err != nil {
			fmt.Println("Error listing institutions:", err)
			return
		}
		for _, institution := range institutions {
			fmt.Println("Name: ", institution.Name, ": institution code =", institution.Code)
		}

		fmt.Println("\nPlease enter in desired institution code: ")
		institutionCode, _ := br.ReadString('\n')
		institutionCode = strings.TrimSuffix(institutionCode, "\n")

		loginCredentials := []*models.Credential{}
		credentials, err := client.ListCredentials(institutionCode)
		if err != nil {
			fmt.Println("Error listing credentials:", err)
			return
		}
		for _, credential := range credentials {
			fmt.Println("\nPlease enter in", credential.Label)
			userInput, _ := br.ReadString('\n')
			userInput = strings.TrimSuffix(userInput, "\n")
			cred := &models.Credential{Guid: credential.Guid, Value: userInput}
			loginCredentials = append(loginCredentials, cred)
		}

		newMember := &models.Member{
			InstitutionCode: institutionCode,
		}
		member, err := client.CreateMember(userGUID, newMember, loginCredentials)
		if err != nil {
			fmt.Println("Error creating member:", err)
			return
		}
		memberGUID = member.Guid
		fmt.Println("\nCreated member:", memberGUID)

		checkJobStatus(client, userGUID, memberGUID)
	} else {
		fmt.Println("\nEnd user must be present to create a new member")
		os.Exit(0)
	}
}

func checkJobStatus(client *client.Client, userGUID string, memberGUID string) {
	fmt.Println("\n2 second delay...")
	time.Sleep(2 * time.Second)

	member, err := client.GetMemberStatus(userGUID, memberGUID)
	if err != nil {
		fmt.Println("Error getting member status:", err)
		return
	}
	status := member.Status

	fmt.Println("\nJOB STATUS:", status)

	if status == "COMPLETED" {
		readAggregationData(client, userGUID, memberGUID)
	} else if status == "HALTED" || status == "FAILED" {
		currentTime := time.Now().UTC().Format("2006-01-02T15:04:05+00:00")

		member, err := client.GetMemberStatus(userGUID, memberGUID)
		if err != nil {
			fmt.Println("Error getting member status:", err)
			return
		}
		lastSuccessTime := member.SuccessfullyAggregatedAt

		// Check if last successful aggregation over 3 days aggregation
		if lastSuccessTime != "nil" && (math.Abs(strToFloat(currentTime[8:10])-strToFloat(lastSuccessTime[8:10])) > 3 || math.Abs(strToFloat(currentTime[5:7])-strToFloat(lastSuccessTime[5:7])) > 0 || math.Abs(strToFloat(currentTime[0:4])-strToFloat(lastSuccessTime[0:4])) > 0) {
			fmt.Println("\nClient should contact MX Support to resolve issue.")
		} else {
			fmt.Println("\nAn update is currently unavailable. Please try again tomorrow")
		}
	} else if status == "CREATED" || status == "UPDATED" || status == "RESUMED" || status == "CONNECTED" || status == "DEGRADED" || status == "DELAYED" || status == "INITIATED" || status == "REQUESTED" || status == "AUTHENTICATED" || status == "RECEIVED" || status == "TRANSFERRED" {
		checkJobStatus(client, userGUID, memberGUID)
	} else if status == "PREVENTED" || status == "DENIED" || status == "IMPAIRED" {
		member, err := client.GetMember(userGUID, memberGUID)
		if err != nil {
			fmt.Println("Error reading member:", err)
			return
		}

		fmt.Println("\nPlease update credentials")
		updatedCredentials := []*models.Credential{}
		br := bufio.NewReader(os.Stdin)
		credentials, err := client.ListCredentials(member.InstitutionCode)
		if err != nil {
			fmt.Println("Error listing credentials:", err)
			return
		}
		for _, credential := range credentials {
			fmt.Println("\nPlease enter in", credential.FieldName)
			userInput, _ := br.ReadString('\n')
			userInput = strings.TrimSuffix(userInput, "\n")
			cred := &models.Credential{Guid: credential.Guid, Value: userInput}
			updatedCredentials = append(updatedCredentials, cred)
		}

		updatedMember := &models.Member{
			Guid: memberGUID,
		}

		member, err = client.UpdateMember(userGUID, updatedMember, updatedCredentials)
		if err != nil {
			fmt.Println("Error updating member:", err)
			return
		}

		checkJobStatus(client, userGUID, memberGUID)
	} else if status == "CHALLENGED" {
		fmt.Println("\nPlease answer the following challenges:")

		br := bufio.NewReader(os.Stdin)
		responses := []*models.Challenge{}

		challenges, err := client.GetMemberChallenges(userGUID, memberGUID)
		if err != nil {
			fmt.Println("Error getting member challenge", err)
			return
		}
		for _, challenge := range challenges {
			fmt.Println(challenge.Label)
			userInput, _ := br.ReadString('\n')
			userInput = strings.TrimSuffix(userInput, "\n")
			response := &models.Challenge{Guid: challenge.Guid, Value: userInput}
			responses = append(responses, response)
		}

		_, err = client.ResumeMember(userGUID, memberGUID, responses)
		if err != nil {
			fmt.Println("Error resuming member:", err)
			return
		}

		checkJobStatus(client, userGUID, memberGUID)
	} else if status == "REJECTED" {
		_, err := client.AggregateMember(userGUID, memberGUID)
		if err != nil {
			fmt.Println("Error aggregating member:", err)
			return
		}

		checkJobStatus(client, userGUID, memberGUID)
	} else if status == "EXPIRED" {
		fmt.Println("\nUser did not answer MFA in time. Please try again tomorrow.")
	} else if status == "LOCKED" {
		fmt.Println("\nUser's account is locked at FI")
	} else if status == "IMPEDED" {
		fmt.Println("\nUser's attention is required at FI website in order for aggregation to complete")
	} else if status == "DISCONTINUED" {
		fmt.Println("\nConnection to institution is no longer available.")
	} else if status == "CLOSED" || status == "DISABLED" {
		fmt.Println("\nAggregation is purposely turned off for this user.")
	} else if status == "TERMINATED" || status == "ABORTED" || status == "STOPPED" || status == "THROTTLED" || status == "SUSPENDED" || status == "ERRORED" {
		if counter > 3 {
			counter = counter + 1
			checkJobStatus(client, userGUID, memberGUID)
		} else {
			fmt.Println("\nAn update is currently unavailable. Please try again tomorrow and contact support if unsuccessful after 3 days.")
			counter = 0
		}
	} else {
		fmt.Println(status)
	}
}

func readAggregationData(client *client.Client, userGUID string, memberGUID string) {
	fmt.Println("\n* Listing all member accounts and transactions *")
	accounts, err := client.ListMemberAccounts(userGUID, memberGUID)
	if err != nil {
		fmt.Println("Error listing member accounts:", err)
		return
	}
	for _, account := range accounts {
		fmt.Println("\nType:", account.Type, "\tName:", account.Name, "\tAvailable Balance:", account.AvailableBalance, "\tAvailable Credit:", account.AvailableCredit)
		fmt.Println("Transactions")
		transactions, accountError := client.ListAccountTransactions(userGUID, account.Guid)
		if err != nil {
			fmt.Println("Error listing account transactions:", accountError)
			return
		}
		for _, transaction := range transactions {
			fmt.Println("\tDate:", transaction.PostedAt, "\tDescription:", transaction.Description, "\tAmount:", transaction.Amount)
		}
	}
}

func strToFloat(value string) float64 {
	float, _ := strconv.ParseFloat(value, 64)
	return float
}
