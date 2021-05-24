package server

import (
	"database/sql"
	uuid "github.com/satori/go.uuid"
	"log"
)

type User struct {
	id       uuid.UUID
	name     string
	objectId uuid.NullUUID
	password string
}

type Users struct {
	items map[string]*User
}

func (o *Users) Count() int {
	return len(o.items)
}

func (o *Users) Load(dsn string) {
	o.items = make(map[string]*User) // warm up

	// ready the connection
	db, err := sql.Open("mysql", dsn)
	if db == nil || err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// fetch from database
	res, err := db.Query("SELECT * FROM user")
	if res == nil || err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	// add items
	for res.Next() {
		var tmp User
		err := res.Scan(&tmp.id, &tmp.name, &tmp.objectId, &tmp.password)
		if err != nil {
			log.Fatal(err)
		}

		o.items[tmp.id.String()] = &tmp
	}
}
