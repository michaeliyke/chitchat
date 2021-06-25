package data

import "testing"

// Delete all threads from database
func ThreadDeleteAll() (err error) {
	statement := "DELETE FROM threads"
	_, err = Db.Exec(statement)
	return
}

func Test_CreateThread(t *testing.T) {
	setUp()
	err := users[0].Create()
	if err != nil {
		t.Error(err, "- Cannot create user")
	}
	conv, err := users[0].CreateThread("My first thread")
	if err != nil {
		t.Error(err, "- Cannot create thread")
	}
	if conv.UserId != users[0].Id {
		t.Error("User not linked with thread")
	}
}
