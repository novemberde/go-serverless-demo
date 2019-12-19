# serverless-go-todo-demo
Build a todo app using golang with serverless framework.

[demo](https://go-todo.judoka.dev/)

## How to start?

1. create ".go-serverless-demo.yml" on root directory. Edit YOUR_BUCKET_NAME.

```yaml
DEPLOYMENT_BUCKET: YOUR_BUCKET_NAME
DYNAMO_REGION: ap-northeast-2
DYNAMO_TABLE_NAME: go-todo
```

2. Create your dynamodb table(go-todo)

3. Run App

- local

```sh
$ go run main.go dev
```

- serverless

```sh
$ export DEPLOYMENT_BUCKET=YOUR_BUCKET_NAME
$ export DYNAMO_REGION=ap-northeast-2
$ export DYNAMO_TABLE_NAME=go-todo
## Build and Deploy
$ make deploy
```

## Test

```sh
$ go test -v $(go list ./... | grep -v vendor) -timeout 15s --count 1 -race -coverprofile=c.out -covermode=atomic
```

## License

[MIT License](/LICENSE)