package main

import (
	"github.com/tidwall/gjson"

	"github.com/forstmeier/backend/graphql"
)

func (dc *dgraphClient) addUsers(ownerID string, ownerIndex int, labIDs, storageIDs []string) ([]string, error) {
	inputs := []map[string]interface{}{}

	start := (len(users) / len(orgs)) * ownerIndex
	end := start + (len(users) / len(orgs))
	if end > len(users) {
		end = len(users) - 1
	}

	for _, user := range users[start:end] {
		orgID := ""
		if user.role == "USER_STORAGE" {
			orgID = randomString(storageIDs)
		} else if user.role == "USER_LAB" {
			orgID = randomString(labIDs)
		} else {
			orgID = ownerID
		}

		owner := map[string]string{
			"id": ownerID,
		}
		org := map[string]string{
			"id": orgID,
		}
		input := map[string]interface{}{
			"owner":     owner,
			"email":     user.email,
			"firstName": user.first,
			"lastName":  user.last,
			"role":      user.role,
			"org":       org,
			"auth0ID":   user.userID,
		}

		inputs = append(inputs, input)
	}

	data, err := dc.SendRequest(graphql.AddUsersMutation, inputs)
	if err != nil {
		return nil, err
	}

	emailValues := gjson.Get(data, "data.addUser.user.#.email").Array()
	emails := []string{}
	for _, email := range emailValues {
		emails = append(emails, email.String())
	}

	return emails, nil
}
