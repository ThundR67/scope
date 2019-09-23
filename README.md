[![Go Report Card](https://goreportcard.com/badge/github.com/SonicRoshan/scope)](https://goreportcard.com/report/github.com/SonicRoshan/scope) [![GoDoc](https://godoc.org/github.com/SonicRoshan/scope?status.svg)](https://godoc.org/github.com/SonicRoshan/scope) [![GoCover](https://gocover.io/_badge/github.com/SonicRoshan/scope)](https://gocover.io/github.com/SonicRoshan/scope)

# Scope
Easily Manage OAuth2 Scopes In Go

## Scope Matching Using Wildcard Strategy
```go
import "github.com/SonicRoshan/scope"

scopeA := "read:user:*"
scopeB := "read:user:username"

doesMatch := scope.MatchScopes(scopeA, scopeB)
```
This strategy will work like this :-
* `users.*` matches `users.read`
* `users.*` matches `users.read.foo`
* `users.read` matches `users.read`
* `users` does not match `users.read`
* `users.read.*` does not match `users.read`
* `users.*.*` does not match `users.read`
* `users.*.*` matches `users.read.own`
* `users.*.*` matches `users.read.own.other`
* `users.read.*` matches `users.read.own`
* `users.read.*` matches `users.read.own.other`
* `users.write.*` does not match `users.read.own`
* `users.*.bar` matches `users.baz.bar`
* `users.*.bar` does not `users.baz.baz.bar`

## Filtering Struct For Read Request
When a client request certain data, this function will eliminate any data in the struct for which the client does not have a read scope.
```go
type user struct {
    username string `readScope:"user:read:username"`
    email string `readScope:"user:read:email"`
}


func main() {
    output := user{username : "Test", email : "Test@Test.com"}
    scopesHeldByClient := []string{"user:read:username"}
    scope.FilterRead(output, scopesHeldByClient)

    // Now output.email will be nil as client does not have scope required to read email field

    output := user{username : "Test", email : "Test@Test.com"}
    scopesHeldByClient := []string{"user:read:*"}
    scope.FilterRead(output, scopesHeldByClient)

    // Now none of the field in output will be nil as client has scopes to read everything in user struct
}

```