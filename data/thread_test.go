package data

import "testing"

// Delete all threads from database
func ThreadDeleteAll() (err error) {
	_, err = Db.Exec("DELETE FROM threads")
	throw(err, "could not delete threads")
	_, err = Db.Exec("ALTER SEQUENCE threads_id_seq RESTART WITH 1")
	throw(err, "sequence could not be reset on threads table")
	return
}

func Test_CreateThread(t *testing.T) {
	setUp()
	err := users[0].Create()
	if err != nil {
		t.Fatal(err, "- Cannot create user")
	}
	conv, err := users[0].CreateThread("My first thread")
	if err != nil {
		t.Fatal(err, "- Cannot create thread")
	}
	if conv.UserId != users[0].Id {
		t.Fatal("User not linked with thread")
	}
}
