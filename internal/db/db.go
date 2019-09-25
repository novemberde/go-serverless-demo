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
	Username  string    `json:"username,omitempty" dynamo:"sk"`
	CreatedAt time.Time `json:"created_at,omitempty" dynamo:"CreatedAt,omitempty"`
}

// Todo ...
type Todo struct {
	Username  string    `json:"username,omitempty" dynamo:"pk"`   // pk
	CreatedAt time.Time `json:"created_at,omitempty" dynamo:"sk"` // sk
	Content   string    `json:"content" dynamo:"Content"`
	UserAgent string    `dynamo:"UserAgent,omitempty"`
	Meta      string    `json:"meta,omitempty" dynamo:"Meta"`
	UpdatedAt time.Time `json:"updated_at" dynamo:"UpdatedAt"`
	DeletedAt time.Time `json:"deleted_at,omitempty" dynamo:"DeletedAt"`
	Checked   bool      `json:"checked" dynamo:"Checked"`
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
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	return db.todo.Put(t).Run()
}

// Update ...
func (db *DB) Update(t *Todo) error {
	q := db.todo.Update("pk", t.Username).Range("sk", t.CreatedAt)

	if t.Content != "" {
		q.Set("Content", t.Content)
	}
	q.SetExpr("Checked=?", t.Checked)

	return q.
		Set("UpdatedAt", time.Now()).
		Run()
}

// Delete ...
func (db *DB) Delete(t *Todo) error {
	return db.todo.Delete("pk", t.Username).Range("sk", t.CreatedAt).
		Run()
}

// Check ...
func (db *DB) Check(t *Todo) error {
	return db.todo.Update("pk", t.Username).Range("sk", t.CreatedAt).
		Set("CheckedAt", time.Now()).
		Run()
}

// Find ...
func (db *DB) Find(username string) ([]Todo, error) {
	var todos []Todo
	err := db.todo.Get("pk", username).All(&todos)

	if err != nil {
		return nil, err
	}
	return todos, nil
}
