package model_test

import (
	"testing"

	"github.com/kotoninja/http-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_Validation(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isvalid bool
	}{
		{name: "Valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isvalid: true,
		},
		{name: "empty Email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isvalid: false,
		},
		{name: "invalid Email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "invalid"
				return u
			},
			isvalid: false,
		},
		{name: "with encrypted password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "EncryptedPassword"
				return u
			},
			isvalid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isvalid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
	// u := model.TestUser(t)
	// assert.NoError(t, u.Validate())
}
func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)

}
