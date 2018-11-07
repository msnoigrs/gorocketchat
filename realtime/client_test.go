package realtime

import (
	"net/url"
	"testing"

	"github.com/msnoigrs/gorocketchat/common_testing"
	"github.com/msnoigrs/gorocketchat/models"
	"github.com/stretchr/testify/assert"
)

var (
	client *Client
)

func getLoggedInClient(t *testing.T) *Client {

	if client == nil {
		c, err := NewClient(&url.URL{Host: common_testing.Host + ":" + common_testing.Port}, true)
		assert.Nil(t, err, "Couldn't create realtime client")

		_, err = c.RegisterUser(&models.UserCredentials{
			Email:    common_testing.GetRandomEmail(),
			Name:     common_testing.GetRandomString(),
			Password: common_testing.GetRandomString()})
		assert.Nil(t, err, "Couldn't register user")

		client = c
	}

	return client
}
