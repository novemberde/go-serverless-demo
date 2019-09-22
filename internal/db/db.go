package handler

import (
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

const usersPartitionKey = "Hl4WRzKE7yvYCz4PFn4O6CFD01lBVL1"

// DB ...
type DB struct {
	todo dynamo.Table
}

// User ...
type User struct {
	PK        string    `dynamo:"pk,omitempty"`
	Username  string    `dynamo:"sk,omitempty"`
	CreatedAt time.Time `dynamo:"CreatedAt,omitempty"`
}

// Todo ...
type Todo struct {
	Username  string    `dynamo:"pk,omitempty"` // pk
	CreatedAt time.Time `dynamo:"sk,omitempty"` // sk

	Content   string    `dynamo:"Content"`
	Meta      string    `dynamo:"Meta,omitempty"`
	UpdatedAt time.Time `dynamo:"UpdatedAt"`
	DeletedAt time.Time `dynamo:"DeletedAt,omitempty"`
	CheckedAt time.Time `dynamo:"CheckedAt,omitempty"`
}

// New ...
func New(region, tableName string) *DB {
	return &DB{
		todo: dynamo.New(session.New(), &aws.Config{Region: aws.String(region)}).Table(tableName),
	}
}

// AddUser ...
func (db *DB) AddUser(username string) error {
	username = strings.Trim(username, " ")
	u := &User{
		PK:        usersPartitionKey,
		CreatedAt: time.Now(),
		Username:  username,
	}
	var exist *User
	err := db.todo.Get("pk", usersPartitionKey).Range("sk", dynamo.Equal, username).One(&exist)

	if err == nil && exist != nil {
		log.Println("exists user. username: ", username)
	}

	return db.todo.Put(u).Run()
}

func (db *DB) deleteUser(username string) error {
	return db.todo.Delete("pk", usersPartitionKey).Range("sk", username).Run()
}

// Create ...
func (db *DB) Create(t *Todo) error {
	return db.todo.Put(t).Run()
}

// Update ...
func (db *DB) Update(t *Todo) error {
	return db.todo.Update("pk", t.Username).Range("sk", t.CreatedAt).
		Set("Content", t.Content).
		Set("UpdatedAt", time.Now()).
		Run()
}

// Delete ...
func (db *DB) Delete(t *Todo) error {
	return db.todo.Update("pk", t.Username).Range("sk", t.CreatedAt).
		Set("DeletedAt", time.Now()).
		Run()
}

// Check ...
func (db *DB) Check(t *Todo) error {
	return db.todo.Update("pk", t.Username).Range("sk", t.CreatedAt).
		Set("CheckedAt", time.Now()).
		Run()
}

// Find ...
func (db *DB) Find() ([]Todo, error) {
	var todos []Todo
	err := db.todo.Scan().All(&todos)

	if err != nil {
		return nil, err
	}
	return todos, nil
}
