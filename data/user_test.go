package data

import (
	"database/sql"
	"testing"
)

func Test_UserCreate(t *testing.T) {
	setUp()
	err := users[0].Create()
	if err != nil {
		t.Fatal(err, "- Cannot create user")
	}
	if users[0].Id == 0 {
		t.Fatalf("No Id, or created_at in user")
	}
	u, err := UserByEmail(users[0].Email)
	if err != nil {
		t.Fatal(err, "- User not created")
	}
	if users[0].Email != u.Email {
		t.Fatal("User retrieved is not the same as the one created")
	}
}

func Test_UserDelete(t *testing.T) {
	setUp()
	err := users[0].Create()
	if err != nil {
		t.Fatal(err, "- Cannot create user")
	}
	err = users[0].Delete()
	if err != nil {
		t.Fatal(err, "- Cannot delete user")
	}
	_, err = UserByEmail(users[0].Email)
	if err != sql.ErrNoRows {
		t.Fatal(err, "- user not deleted")
	}
}

func Test_UserUpdate(t *testing.T) {
	setUp()
	err := users[0].Create()
	if err != nil {
		t.Fatal(err, "- cannot create user")
	}
	users[0].Name = "Random User"
	err = users[0].Update()
	if err != nil {
		t.Fatal(err, "- cannot update user")
	}
	u, err := UserByEmail(users[0].Email)
	if err != nil {
		t.Fatal(err, "- cannot get user")
	}
	if u.Name != "Random User" {
		t.Fatal(err, "- user no updated")
	}
}

func Test_UserByUUID(t *testing.T) {
	setUp()
	err := users[0].Create()
	if err != nil {
		t.Fatal(err, "- cannot create user")
	}
	u, err := UserByUUID(users[0].Uuid)
	if err != nil {
		t.Fatal(err, "- user not created")
	}
	if users[0].Email != u.Email {
		t.Fatalf("user retrieved is not the same as the one created")
	}
}

func Test_Users(t *testing.T) {
	setUp()
	for _, user := range users {
		err := user.Create()
		if err != nil {
			t.Fatal(err, "- cannot create user")
		}
	}
	u, err := Users()
	if err != nil {
		t.Fatal(err, "- cannot retrieve users")
	}
	if len(u) != len(users) {
		t.Fatal(err, "- wrong number of users retrieved")
	}
	if u[0].Email != users[0].Email {
		t.Fatal(u[0], users[0], "- wrong user retrieved")
	}
}

func Test_CreateSession(t *testing.T) {
	setUp()
	err := users[0].Create()
	if err != nil {
		t.Fatal(err, "- Cannot create user")
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Fatal(err, "- Cannot create session")
	}
	if session.UserId != users[0].Id {
		t.Fatal("User not linked with session")
	}
}

func Test_GetSession(t *testing.T) {
	setUp()
	err := users[4].Create()
	if err != nil {
		t.Fatal(err, "- Cannot create user")
	}
	Session, err := users[4].CreateSession()
	if err != nil {
		t.Fatal(err, "- Cannot create session")
	}
	s, err := users[4].Session()
	if err != nil {
		t.Fatal(err, "- Cannot get session")
	}
	if s.Id == 0 {
		t.Fatal("No session retrieved")
	}
	if s.Id != Session.Id {
		t.Fatal("Different session retrieved")
	}
}

func Test_CheckValidSession(t *testing.T) {
	setUp()
	err := users[0].Create()
	if err != nil {
		t.Fatal(err, "- Cannot create user")
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Fatal(err, "- Cannot create session")
	}
	uuid := session.Uuid
	s := Session{Uuid: uuid}
	valid, err := s.Check()
	if err != nil {
		t.Fatal(err, "- Cannot check session")
	}
	if valid != true {
		t.Fatal(err, "- Session is not valid")
	}
}

func Test_CheckInvalidSession(t *testing.T) {
	setUp()
	s := Session{Uuid: "123"}
	valid, err := s.Check()
	if err == nil {
		t.Fatal(err, "- Session is not valid but is validated")
	}
	if valid == true {
		t.Fatal(err, "- Session is valid")
	}
}

func Test_DeleteSession(t *testing.T) {
	setUp()
	err := users[0].Create()
	if err != nil {
		t.Fatal(err, "- Cannot create user")
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Fatal(err, "- Cannot create session")
	}
	err = session.DeleteByUUID()
	if err != nil {
		t.Fatal(err, "- Cannot delete session")
	}
	s := Session{Uuid: session.Uuid}
	valid, err := s.Check()
	if err == nil {
		t.Fatal(err, "- Session is valid even though deleted")
	}
	if valid == true {
		t.Fatal(err, "- Session is not deleted")
	}
}
