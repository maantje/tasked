package controller

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"tasked/database"
	"tasked/database/migration"
	"tasked/test"
	"testing"
)

func TestCreateUser(t *testing.T) {
	database.InitTest()
	migration.Migrate()

	c, rec := test.PostJSON("/users", `{
		"name": "Jon Snow",
		"email": "jon@labstack.com",
		"password": "secret"
	}`)

	uc := &UserController{}

	// Assertions
	if assert.NoError(t, uc.Create(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		test.JSONEq(t, rec.Body.String(), `{
				"id": 1,
				"created_at": "<<PRESENCE>>",
				"updated_at": "<<PRESENCE>>",
				"name": "Jon Snow",
				"email": "jon@labstack.com",
				"email_verified_at": null
			}`,
		)
	}
}
