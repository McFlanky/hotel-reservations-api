package api

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/McFlanky/hotel-reservations-api/types"
	"github.com/gofiber/fiber/v2"
)

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := NewUserHandler(tdb.User)
	app.Post("/", userHandler.HandlePostUser)

	params := types.CreateUserParams{
		Email:     "some@foo.com",
		FirstName: "foo",
		LastName:  "bar",
		Password:  "slgnslglsgljljnbl",
	}
	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	var user types.User
	json.NewDecoder(resp.Body).Decode(&user)
	if len(user.ID) == 0 {
		t.Errorf("expected a user ID to be set")
	}
	if len(user.EncryptedPassword) > 0 {
		t.Errorf("expected encrypted password not to be included in json response")
	}
	if user.FirstName != params.FirstName {
		t.Errorf("expected username: %s but got: %s", params.FirstName, user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("expected lastname: %s but got: %s", params.LastName, user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("expected email: %s but got: %s", params.Email, user.Email)
	}
}
