package ghclient

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// ParseHeaders is used to return the EventID and GitHubEvent from request headers
func ParseHeaders(r *http.Request) (string, string) {
	EventID := r.Header.Get("X-GitHub-Delivery")
	GitHubEvent := r.Header.Get("X-GitHub-Event")
	return EventID, GitHubEvent
}

// IsValidSignature validates the message body with the checksum sent by GitHub
func IsValidSignature(r *http.Request, key string) bool {
	gotHash := strings.SplitN(r.Header.Get("X-Hub-Signature"), "=", 2)
	if gotHash[0] != "sha1" {
		log.Panicf("Checksum is invalid")
		return false
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panicf("Cannot read the request body: %s", err)
		return false
	}

	hash := hmac.New(sha1.New, []byte(key))
	if _, err := hash.Write(b); err != nil {
		log.Panicf("Cannot compute the HMAC for request: %s", err)
		return false
	}

	expectedHash := hex.EncodeToString(hash.Sum(nil))
	isValid := gotHash[1] == expectedHash
	if isValid == false {
		log.Panicf("Invalid Hash: %s", expectedHash)
	}
	return isValid
}
