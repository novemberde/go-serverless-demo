package db

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DB", func() {
	const (
		tableName = "go-todo-test"

		username = "username"
		content  = "Hello world"
	)

	var d *DB

	BeforeEach(func() {
		d = New(&aws.Config{
			Endpoint: aws.String("http://localhost:8000"),
			Region:   aws.String("ap-northeast-2"),
		})

		d.db.Table(tableName).DeleteTable().Run()

		err := d.CreateTable(tableName, new(Todo))
		if err != nil {
			log.Panicln("err:", err)
		}

		d.SetTable(tableName)
	})

	AfterEach(func() {
		d.db.Table(tableName).DeleteTable().Run()
	})

	Describe("Create", func() {
		var todo *Todo

		Context("with valid todo", func() {
			BeforeEach(func() {
				todo = &Todo{
					Username: username,
					Content:  content,
				}
			})

			It("creates a document", func() {
				originTodos, err := d.Find(username)
				if err != nil {
					log.Panicln(err)
				}

				err = d.Create(todo)
				if err != nil {
					log.Panicln(err)
				}

				todos, err := d.Find(username)
				if err != nil {
					log.Panicln(err)
				}

				Expect(len(todos) - len(originTodos)).To(Equal(1))
			})
		})

		Context("without username", func() {
			BeforeEach(func() {
				todo = &Todo{
					Content: content,
				}
			})

			It("occurs error", func() {
				err := d.Create(todo)

				Expect(err.Error()).
					To(ContainSubstring("One of the required keys was not given a value"))
			})
		})
	})

	Describe("Find", func() {
		Context("with existing username", func() {
			todo := &Todo{
				Username: username,
				Content:  content,
			}

			BeforeEach(func() {
				err := d.Create(todo)
				if err != nil {
					log.Panicln(err)
				}
			})

			It("returns todos", func() {
				todos, err := d.Find(username)
				if err != nil {
					log.Panicln(err)
				}

				Expect(todos[0].Username).To(Equal(todo.Username))
				Expect(todos[0].Content).To(Equal(todo.Content))
			})
		})

		Context("with not existing user", func() {
			It("returns empty todos", func() {
				todos, err := d.Find("NOT_EXISTING")
				if err != nil {
					log.Panicln(err)
				}

				Expect(todos).To(BeEmpty())
			})
		})
	})

	Describe("Update", func() {
		var todo Todo

		BeforeEach(func() {
			err := d.Create(&Todo{
				Username: username,
				Content:  content,
			})
			if err != nil {
				log.Panicln(err)
			}

			todos, err := d.Find(username)
			if err != nil {
				log.Panicln(err)
			}

			todo = todos[0]
		})

		Context("when content change", func() {
			It("updates todo", func() {
				todo.Content = "New content"
				todo.Checked = true
				err := d.Update(&todo)
				if err != nil {
					log.Panicln(err)
				}

				todos, err := d.Find(username)
				if err != nil {
					log.Panicln(err)
				}

				updated := todos[0]

				Expect(updated.Content).To(Equal(todo.Content))
				Expect(updated.Checked).To(Equal(todo.Checked))
			})
		})

		Context("when content is empty", func() {
			It("nothing happens", func() {
				todo.Content = ""
				err := d.Update(&todo)
				if err != nil {
					log.Panicln(err)
				}

				todos, err := d.Find(username)
				if err != nil {
					log.Panicln(err)
				}

				updated := todos[0]

				Expect(updated.Content).To(Equal(content))
				Expect(updated.Checked).To(Equal(todo.Checked))
			})
		})
	})

	Describe("Delete", func() {
		var todo Todo

		Context("with existing Todo", func() {
			BeforeEach(func() {
				err := d.Create(&Todo{
					Username: username,
					Content:  content,
				})
				if err != nil {
					log.Panicln(err)
				}

				todos, err := d.Find(username)
				if err != nil {
					log.Panicln(err)
				}

				todo = todos[0]
			})

			It("deletes todo", func() {
				err := d.Delete(&todo)
				if err != nil {
					log.Panicln(err)
				}

				todos, err := d.Find(username)
				if err != nil {
					log.Panicln(err)
				}

				Expect(todos).To(BeEmpty())
			})
		})

		Context("with not existing Todo", func() {
			BeforeEach(func() {
				todo = Todo{
					Username: username,
					Content:  content,
				}
			})

			It("nothing happens", func() {
				err := d.Delete(&todo)

				Expect(err).NotTo(HaveOccurred())
			})
		})
	})
})
