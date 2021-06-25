package data

import (
	"database/sql"
	"testing"
)

func Test_DeleteUsers(t *testing.T) {
	UserDeleteAll()
	_, err := getLastId()
	if err != nil {
		t.Error("No row was found")
	}
}

func Test_CreateUser(t *testing.T) {
	err := users[0].Create()
	if err != nil {
		t.Error(err, "- cannot create user")
	}
}

func Test_UserCreate(t *testing.T) {
	setUp()
	err := users[0].Create()
	if err != nil {
		t.Error(err, "- Cannot create user")
	}
	if users[0].Id == 0 {
		t.Errorf("No Id, or created_at in user")
	}
	u, err := UserByEmail(users[0].Email)
	if err != nil {
		t.Error(err, "- User not created")
	}
	if users[0].Email != u.Email {
		t.Errorf("User retrieved is not the same as the one created")
	}
}

func Test_UserDelete(t *testing.T) {
	setUp()
	err := users[0].Create()
	if err != nil {
		t.Error(err, "- Cannot create user")
	}
	err = users[0].Delete()
	if err != nil {
		t.Error(err, "- Cannot delete user")
	}
	_, err = UserByEmail(users[0].Email)
	if err != sql.ErrNoRows {
		t.Error(err, "- user not deleted")
	}
}

func Test_UserUpdate(t *testing.T) {
	setUp()
	err := users[0].Create()
	if err != nil {
		t.Error(err, "- Cannot create user")
	}
	users[0].Name = "Random User"
	err = users[0].Update()
	if err != nil {
		t.Error(err, "- Cannot update user")
	}
	u, err := UserByEmail(users[0].Email)
	if err != nil {
		t.Error(err, "- Cannot get user")
	}
	if u.Name != "Random User" {
		t.Error(err, "- User no updated")
	}
}

func Test_UserByUUID(t *testing.T) {
	setUp()
	err := users[0].Create()
	if err != nil {
		t.Error(err, "- Cannot create user")
	}
	u, err := UserByUUID(users[0].Uuid)
	if err != nil {
		t.Error(err, "- User not created")
	}
	if users[0].Email != u.Email {
		t.Errorf("User retrieved is not the same as the one created")
	}
}

func Test_Users(t *testing.T) {
	setUp()
	for _, user := range users {
		err := user.Create()
		if err != nil {
			t.Error(err, "- Cannot create user")
		}
	}
	u, err := Users()
	if err != nil {
		t.Error(err, "- Cannot retrieve users")
	}
	if len(u) != 2 {
		t.Error(err, "- Wrong of users retrieved")
	}
	if u[0].Email != users[0].Email {
		t.Error(u[0], users[0], "- Wrong user retrieved")
	}
}

func Test_CreateSession(t *testing.T) {
	setUp()
	err := users[0].Create()
	if err != nil {
		t.Error(err, "- Cannot create user")
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Error(err, "- Cannot create session")
	}
	if session.UserId != users[0].Id {
		t.Error("User not linked with session")
	}
}

func Test_GetSession(t *testing.T) {
	setUp()
	err := users[0].Create()
	if err != nil {
		t.Error(err, "- Cannot create user")
	}
	Session, err := users[0].CreateSession()
	if err != nil {
		t.Error(err, "- Cannot create session")
	}
	s, err := users[0].Session()
	if err != nil {
		t.Error(err, "- Cannot get session")
	}
	if s.Id == 0 {
		t.Error("No session retrieved")
	}
	if s.Id != Session.Id {
		t.Error("Different session retrieved")
	}
}

func Test_CheckValidSession(t *testing.T) {
	setUp()
	err := users[0].Create()
	if err != nil {
		t.Error(err, "- Cannot create user")
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Error(err, "- Cannot create session")
	}
	uuid := session.Uuid
	s := Session{Uuid: uuid}
	valid, err := s.Check()
	if err != nil {
		t.Error(err, "- Cannot check session")
	}
	if valid != true {
		t.Error(err, "- Session is not valid")
	}
}

func Test_CheckInvalidSession(t *testing.T) {
	setUp()
	s := Session{Uuid: "123"}
	valid, err := s.Check()
	if err == nil {
		t.Error(err, "- Session is not valid but is validated")
	}
	if valid == true {
		t.Error(err, "- Session is valid")
	}
}

func Test_DeleteSession(t *testing.T) {
	setUp()
	err := users[0].Create()
	if err != nil {
		t.Error(err, "- Cannot create user")
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Error(err, "- Cannot create session")
	}
	err = session.DeleteByUUID()
	if err != nil {
		t.Error(err, "- Cannot delete session")
	}
	s := Session{Uuid: session.Uuid}
	valid, err := s.Check()
	if err == nil {
		t.Error(err, "- Session is valid even though deleted")
	}
	if valid == true {
		t.Error(err, "- Session is not deleted")
	}
}
