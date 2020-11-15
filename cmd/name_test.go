package cmd

import (
	"os/user"
	"testing"
)

func TestgetUserName(t *testing.T) {
	// TODO: Find a way to mock the User.user object instead.
	testUser, testErr := user.Current()
	if testErr != nil {
		panic(testErr)
	}
	getUserNameResult := getUserName(testUser)
	if getUserNameResult == "" {
		t.Errorf("getUserName() returned empty string!")
	}
}

func TestgetName(t *testing.T) {
	// TODO: Find a way to mock the User.user object instead.
	testUser, testErr := user.Current()
	if testErr != nil {
		panic(testErr)
	}
	getUserNameResult := getName(testUser)
	if getUserNameResult == "" {
		t.Errorf("getUserName() returned empty string!")
	}
}
