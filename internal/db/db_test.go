package handler

import (
	"testing"
	"time"

	"github.com/guregu/dynamo"
)

func TestDB_Create(t *testing.T) {
	db := New("ap-northeast-2", "go-todo")
	type fields struct {
		todo dynamo.Table
	}
	type args struct {
		t *Todo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"ok",
			fields{
				todo: db.todo,
			},
			args{
				t: &Todo{
					Username:  "Test",
					CreatedAt: time.Now(),
					Content:   "Hello World",
					UpdatedAt: time.Now(),
				},
			},
			false,
		},
		{
			"invalid Username",
			fields{
				todo: db.todo,
			},
			args{
				t: &Todo{
					CreatedAt: time.Now(),
					Content:   "Hello World",
					UpdatedAt: time.Now(),
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				todo: tt.fields.todo,
			}
			if err := db.Create(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("DB.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_AddUser(t *testing.T) {
	db := New("ap-northeast-2", "go-todo")
	username := "Test"
	type fields struct {
		todo dynamo.Table
	}
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"ok",
			fields{
				todo: db.todo,
			},
			args{
				username: username,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				todo: tt.fields.todo,
			}
			if err := db.AddUser(tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("DB.AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	db.deleteUser(username)
}

func TestDB_Update(t *testing.T) {
	db := New("ap-northeast-2", "go-todo")
	todo := &Todo{
		Username: "Test",
		Content:  "Hello World Update",
	}
	db.Create(todo)
	type fields struct {
		todo dynamo.Table
	}
	type args struct {
		t *Todo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"ok",
			fields{
				todo: db.todo,
			},
			args{
				t: &Todo{
					Username:  todo.Username,
					CreatedAt: todo.CreatedAt,
					Checked:   true,
				},
			},
			false,
		},
		{
			"ok2",
			fields{
				todo: db.todo,
			},
			args{
				t: &Todo{
					Username:  todo.Username,
					CreatedAt: todo.CreatedAt,
					Checked:   false,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				todo: tt.fields.todo,
			}
			if err := db.Update(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("DB.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
