package main

import (
	"fmt"
	"github.com/hunterlong/statup/types"
	"time"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

var (
	Database sqlbuilder.Database
)

type Communication types.Communication

func MakeThenDelete() {
	// Create new communication pointer
	newComm := New()

	// Insert the new communication object into the database
	newComm.Create()
	fmt.Printf("Inserted a new Communication method into database, ID: %v\n", newComm.Id)

	// Read the new communication that was inserted into database
	comm := Read(newComm.Id)
	fmt.Printf("Select the communication from database, Method: %v\n", comm.Method)

	// Update the communication with some variables, and then update back to database
	comm.ApiKey = "0x0x0x00x0x0"
	comm.ApiSecret = "727SECRETS0x9199F9F"
	comm.Update()
	fmt.Printf("Updated the communication to database, Method: %v, API: %v\n", comm.Method, comm.ApiKey)

	// This is just an example, now its being deleted.
	comm.Delete()
	fmt.Printf("Deleted the communication object from database.\n\n")
}

// A basic function that returns the database column "communications"
// this could be changed to use any table from the database
func CommunicationTable() db.Collection {
	return Database.Collection("communication")
}

// Create new pointer for a communication
// Methods should be unique for each plugin used!
func New() *Communication {
	comm := &Communication{
		Method:    "Ping",
		CreatedAt: time.Now(),
	}
	return comm
}

// Insert a new communication into database
// once inserted, return the Communication
func (c *Communication) Create() *Communication {
	uuid, err := CommunicationTable().Insert(c)
	if err != nil {
		panic(err)
	}
	c.Id = uuid.(int64)
	return c
}

// Read from database and return a Communication object
func Read(id int64) *Communication {
	var comm *Communication
	col := CommunicationTable().Find("id", id)
	err := col.One(&comm)
	if err != nil {
		panic(err)
	}
	return comm
}

// Update a communication to database
func (c *Communication) Update() *Communication {
	col := CommunicationTable().Find("id", c.Id)
	err := col.Update(c)
	if err != nil {
		panic(err)
	}
	return c
}

// Delete a communication from database
func (c *Communication) Delete() *Communication {
	col := CommunicationTable().Find("id", c.Id)
	err := col.Delete()
	if err != nil {
		panic(err)
	}
	return c
}
