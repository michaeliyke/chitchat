package data

import "log"

//nolint:unused,deadcode
//nolint:staticcheck
var _ = throw

func throw(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

var _ = warn

func warn(err error, msg string) {
	if err != nil {
		log.Println(msg, err)
	}
}
