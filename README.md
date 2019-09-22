# serverless-go-todo-demo
Build a todo app using golang with serverless framework.

## Test

```sh
$ go test -v $(go list ./... | grep -v vendor) -timeout 15s --count 1 -race -coverprofile=c.out -covermode=atomic
```