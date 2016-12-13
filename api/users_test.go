package api

import (
	"encoding/json"
	"fmt"
	"github.com/bigokro/gruff-server/gruff"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"testing"
)

func TestListUsers(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createUser("test1", "test1@test1.com")
	u2 := createUser("test2", "test2@test2.com")
	TESTDB.Create(&u1)
	TESTDB.Create(&u2)

	expectedResults, _ := json.Marshal([]gruff.User{u1, u2})

	r.GET("/api/users")
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestListUsersPagination(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createUser("test1", "test1@test1.com")
	u2 := createUser("test2", "test2@test2.com")
	TESTDB.Create(&u1)
	TESTDB.Create(&u2)

	r.GET("/api/users?start=0&limit=25")
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestGetUsers(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createUser("test1", "test1@test1.com")
	TESTDB.Create(&u1)

	expectedResults, _ := json.Marshal(u1)

	r.GET(fmt.Sprintf("/api/users/%d", u1.ID))
	res, _ := r.Run(Router())
	assert.Equal(t, string(expectedResults), res.Body.String())
	assert.Equal(t, http.StatusOK, res.Status())
}

func TestCreateUsers(t *testing.T) {
	setup()
	defer teardown()

	r := New(nil)

	u1 := createUser("test1", "test1@test1.com")
	TESTDB.Create(&u1)

	u2 := createUser("test1", "test1@test1.com")

	r.POST("/api/users")
	r.SetBody(u2)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusCreated, res.Status())
}

func TestUpdateUsers(t *testing.T) {
	setup()
	defer teardown()

	r := New(Token)

	u1 := createUser("test1", "test1@test1.com")
	TESTDB.Create(&u1)
	// createUsersElastic(u1.ID, u1)

	r.PUT(fmt.Sprintf("/api/users/%d", u1.ID))
	r.SetBody(u1)
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusAccepted, res.Status())
}

func TestDeleteUsers(t *testing.T) {
	setup()
	defer teardown()
	r := New(Token)

	u1 := createUser("test1", "test1@test1.com")
	TESTDB.Create(&u1)
	// createUsersElastic(u1.ID, u1)

	r.DELETE(fmt.Sprintf("/api/users/%d", u1.ID))
	res, _ := r.Run(Router())
	assert.Equal(t, http.StatusOK, res.Status())
}

func createUser(name string, email string) gruff.User {
	u := gruff.User{
		Name:     name,
		Email:    email,
		Password: "123456",
	}
	password := u.Password
	u.Password = ""
	u.HashedPassword, _ = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return u
}
