## Simple http server using golang

```go
type User struct {
	ID        int
	FirstName string
	LastName  string
}
```

## Docker

### Build Using
`docker build -t go_simple_server .`

### Run using
`docker run --publish 3000:3000 --name simple_server --rm go_simple_server`