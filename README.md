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