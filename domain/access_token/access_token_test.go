package access_token

import (
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	if expirationTime != 24 {
		t.Error("expiration time should be 24 hours")
	}
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	if at.IsExpired() {
		t.Error("brand new access toekn should not be expired")
	}

	if at.AccessToken != "" {
		t.Error("new access toekn should not have defined access token id")
	}

	if at.UserId > 0 {
		t.Error("new access toekn should not have an associateted user id ")

	}
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	if !at.IsExpired() {
		t.Error("empty access token should be expired by default")
	}

	at.Expries = time.Now().UTC().Add(3 * time.Hour).Unix()

	if at.IsExpired() {
		t.Error("access token created three hours from now should NOT be expired by default")
	}
}
