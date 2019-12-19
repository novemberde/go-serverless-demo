package db

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// DB ...
type DB struct {
	db   *dynamo.DB
	todo dynamo.Table
}

// Todo ...
type Todo struct {
	Username  string    `json:"username,omitempty" dynamo:"pk,hash"`
	CreatedAt time.Time `json:"created_at,omitempty" dynamo:"sk,range"`
	Content   string    `json:"content" dynamo:"Content"`
	UserAgent string    `json:"user_agent" dynamo:"UserAgent,omitempty"`
	Meta      string    `json:"meta,omitempty" dynamo:"Meta"`
	UpdatedAt time.Time `json:"updated_at" dynamo:"UpdatedAt"`
	DeletedAt time.Time `json:"deleted_at,omitempty" dynamo:"DeletedAt"`
	Checked   bool      `json:"checked" dynamo:"Checked"`
}

// New ...
func New(config *aws.Config) *DB {
	return &DB{db: dynamo.New(session.New(), config)}
}

// SetTable ...
func (d *DB) SetTable(name string) {
	d.todo = d.db.Table(name)
}

// CreateTable ...
func (d *DB) CreateTable(name string, from interface{}) error {
	return d.db.CreateTable(name, from).Run()
}

// Create ...
func (d *DB) Create(t *Todo) error {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	return d.todo.Put(t).Run()
}

// Update ...
func (d *DB) Update(t *Todo) error {
	q := d.todo.
		Update("pk", t.Username).
		Range("sk", t.CreatedAt)

	if t.Content != "" {
		q.Set("Content", t.Content)
	}
	q.SetExpr("Checked=?", t.Checked)
	q.Set("UpdatedAt", time.Now())

	return q.Run()
}

// Delete ...
func (d *DB) Delete(t *Todo) error {
	return d.todo.
		Delete("pk", t.Username).
		Range("sk", t.CreatedAt).
		Run()
}

// Find ...
func (d *DB) Find(username string) (todos []Todo, err error) {
	err = d.todo.
		Get("pk", username).
		All(&todos)
	return
}
