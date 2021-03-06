package test

import (
	"net/http"
	"testing"

	"github.com/golangci/golangci-api/test/sharedtest"
)

func TestGithubLoginFirstTime(t *testing.T) {
	u := sharedtest.Login(t)
	u.E.PUT("/v1/auth/unlink").Expect().Status(http.StatusOK)

	// it's guaranteed first time login
	sharedtest.Login(t)
}

func TestGithubLoginNotFirstTime(t *testing.T) {
	sharedtest.Login(t)
	sharedtest.Login(t)
}

func TestLoginWithAnotherLogin(t *testing.T) {
	u := sharedtest.Login(t)
	u.A.Equal("golangci", u.GithubLogin)

	defer func(prevProfileHandler http.HandlerFunc) {
		sharedtest.FakeGithubCfg.ProfileHandler = prevProfileHandler
	}(sharedtest.FakeGithubCfg.ProfileHandler)

	wasSent := false
	sharedtest.FakeGithubCfg.ProfileHandler = func(w http.ResponseWriter, r *http.Request) {
		ret := map[string]interface{}{
			"login":      "AnotherLogin",
			"email":      "Another_Email@golangci.com",
			"id":         1, // the same github user id
			"avatar_url": "another Avatar",
			"name":       "Another Name",
		}

		sharedtest.SendJSON(w, ret)
		wasSent = true
	}

	u2 := sharedtest.Login(t)
	u2.A.True(wasSent)
	u2.A.Equal("AnotherLogin", u2.GithubLogin)
	u2.A.Equal("another_email@golangci.com", u2.Email)
	u2.A.Equal("another Avatar", u2.AvatarURL)
	u2.A.Equal("Another Name", u2.Name)
	u.A.Equal(u.ID, u2.ID) // logined into the same user
}
