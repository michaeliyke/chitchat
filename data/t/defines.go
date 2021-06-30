package main

import (
	"database/sql"
	"errors"
	"log"

	"github.com/michaeliyke/chitchat/data"
)

type User struct {
	Id   int
	Name string
}

var Db = data.Db

func init() {
}

var _ = checkQuery

func create() {
	data.Usrs[5].Create()
}

func checkQuery(err error) bool {
	var err_ error
	if err == sql.ErrNoRows {
		err_ = errors.New("sql error")
	}
	return check(err_)
}

var _ = verify

func verify(err error) bool {
	return check(err)
}

var _ = check

func check(err error) bool {
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

var _ = ensure

func ensure(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
