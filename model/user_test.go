package model

import (
	"testing"

	"github.com/prakashsingha/go-webapp/helper"
)

func TestLoginHash(t *testing.T) {
	testDB := new(mockDB)
	testDB.returnedRow = &mockRow{}
	db = testDB

	email := "test@email.com"
	password := "testpassword"
	Login(email, password)

	pwd := helper.Hasher(passwordSalt, email, password)

	if testDB.lastArgs[1] != pwd {
		t.Errorf("Failed to send the correct hashed password")
	}

}
