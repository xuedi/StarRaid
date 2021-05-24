package server

import (
	"database/sql"
	uuid "github.com/satori/go.uuid"
	"log"
	"math/rand"
	"time"
)

type Object struct {
	id                uuid.UUID
	name              string
	location          point
	visualRange       int64
	neighbourActivity time.Duration // 1000=lowest ... 10=highest
	neighbourUpdate   time.Time
	neighbours        map[string]neighbour
}

func (o *Object) updateNeighbours(items map[string]*Object) {
	o.neighbours = make(map[string]neighbour)
	for key, item := range items {
		if o.id == item.id { // no self check
			continue
		}

		newNeighbour := neighbour{}
		newNeighbour.id = item.id
		newNeighbour.distance = o.location.Distance(item.location)
		newNeighbour.setVisibilityTo(item.visualRange)

		o.neighbours[key] = newNeighbour
	}
}

type Objects struct {
	items map[string]*Object
}

func (o *Objects) Count() int {
	return len(o.items)
}

func (o *Objects) Tick() {
	for key, item := range o.items {
		if time.Now().After(item.neighbourUpdate) {
			rand.Seed(time.Now().UnixNano())
			wait := (10 + time.Duration(rand.Intn(10))) * time.Second
			o.items[key].neighbourUpdate = time.Now().Add(wait)
			o.items[key].updateNeighbours(o.items)
			//fmt.Println("Update", o.items[key].name)
		}
	}
}

func (o *Objects) Load(dsn string) {
	o.items = make(map[string]*Object) // warm up

	// ready the connection
	db, err := sql.Open("mysql", dsn)
	if db == nil || err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// fetch from database
	res, err := db.Query("SELECT * FROM objects")
	if res == nil || err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	// add items
	for res.Next() {
		var tmp Object
		err := res.Scan(&tmp.id, &tmp.name, &tmp.location.x, &tmp.location.y)
		if err != nil {
			log.Fatal(err)
		}

		tmp.neighbourActivity = 100 * time.Millisecond
		tmp.neighbourUpdate = time.Now().Add(tmp.neighbourActivity)
		tmp.neighbours = make(map[string]neighbour)
		tmp.visualRange = 400
		o.items[tmp.id.String()] = &tmp
	}
}
