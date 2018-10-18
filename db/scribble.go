package db

import (
	scribble "github.com/nanobox-io/golang-scribble"
)

var client *scribble.Driver

func SetScribble(dir string, option *scribble.Options) {
	db, err := scribble.New(dir, nil)
	if err != nil {
		panic(err)
	}
	client = db
}

func GetClient() *scribble.Driver {
	if client == nil {
		panic("scribble isn't set")
	}
	return client
}
