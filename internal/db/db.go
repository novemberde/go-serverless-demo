package handler

import (
	"crypto/sha1"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// DB ...
type DB struct {
	todo dynamo.Table
}

// User ...
type User struct {
	ID        string    `dynamo:"pk,hash"`
	Username  string    `json:"username,omitempty" dynamo:"sk,range"`
	CreatedAt time.Time `json:"created_at,omitempty" dynamo:"CreatedAt"`
}

func pk(username string) string {
	h := sha1.New()
	h.Write([]byte(username))
	return string(h.Sum(nil))
}

// Todo ...
type Todo struct {
	Username  string    `json:"username,omitempty" dynamo:"pk,hash"`    // pk
	CreatedAt time.Time `json:"created_at,omitempty" dynamo:"sk,range"` // sk
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
	pk := pk(username)
	u := &User{
		ID:        pk,
		CreatedAt: time.Now(),
		Username:  username,
	}
	var exist *User
	err := db.todo.Get("pk", pk).Range("sk", dynamo.Equal, username).One(&exist)

	if err == nil && exist != nil {
		log.Println("exists user. username: ", username)
	}

	return db.todo.Put(u).Run()
}

func (db *DB) deleteUser(username string) error {
	return db.todo.Delete("pk", pk(username)).Range("sk", username).Run()
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
	return db.todo.Delete("pk", t.Username).Range("sk", t.CreatedAt).Run()
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
