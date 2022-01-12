package toastcloud

import (
	"log"
	"strings"

	"github.com/toastate/toastate-sdk-go/internal/apiclient"
)

type Session struct {
	client *apiclient.Client
}

func NewSession() *Session {
	return &Session{
		client: apiclient.NewClient("api.cloud.toastate.com", "v1"),
	}
}

func (sess *Session) SetAuth(auth string) *Session {
	if strings.HasPrefix(auth, "sess_") {
		sess.client = sess.client.SetAuthToken(auth)
	} else {
		log.Fatal("invalid authentication")
	}

	// TODO: API Keys

	return sess
}
