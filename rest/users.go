package rest

import (
	"bytes"
	"encoding/json"

	"github.com/msnoigrs/gorocketchat/models"
)

type logoutResponse struct {
	Status
	Data struct {
		Message string `json:"message"`
	} `json:"data"`
}

type logonResponse struct {
	Status
	Data struct {
		Token  string `json:"authToken"`
		UserID string `json:"userID"`
	} `json:"data"`
}

// Login a user. The Email and the Password are mandatory. The auth token of the user is stored in the Client instance.
//
// https://rocket.chat/docs/developer-guides/rest-api/authentication/login
func (c *Client) Login(credentials *models.UserCredentials) error {
	if c.auth != nil {
		return nil
	}

	if credentials.ID != "" && credentials.Token != "" {
		c.auth = &authInfo{id: credentials.ID, token: credentials.Token}
		return nil
	}

	response := new(logonResponse)
	data := map[string]interface{} {
		"user":credentials.Email,
		"password":credentials.Password,
	}
	jsondata, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := c.Post("login", bytes.NewBuffer(jsondata), response); err != nil {
		return err
	}

	c.auth = &authInfo{id: response.Data.UserID, token: response.Data.Token}
	credentials.ID, credentials.Token = response.Data.UserID, response.Data.Token
	return nil
}

// Logout a user. The function returns the response message of the server.
//
// https://rocket.chat/docs/developer-guides/rest-api/authentication/logout
func (c *Client) Logout() (string, error) {

	if c.auth == nil {
		return "Was not logged in", nil
	}

	response := new(logoutResponse)
	if err := c.Get("logout", nil, response); err != nil {
		return "", err
	}

	return response.Data.Message, nil
}
