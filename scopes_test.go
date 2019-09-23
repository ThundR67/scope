package scopes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestScopesMatching tests the MatchScope Function
func TestScopeMatching(t *testing.T) {
	testingData := []interface{}{
		//ScopeB, ScopeA, bool value showing should ScopeA Match With ScopeB
		[]interface{}{"users:*", "users:read", true},
		[]interface{}{"users:*", "users:read:foo", true},
		[]interface{}{"users:read", "users:read", true},
		[]interface{}{"users", "users:read", false},
		[]interface{}{"users:read:*", "users:read", false},
		[]interface{}{"users:*:*", "users:read", false},
		[]interface{}{"users:*:*", "users:read:own", true},
		[]interface{}{"users:*:*", "users:read:own:other", true},
		[]interface{}{"users:read:*", "users:read:own", true},
		[]interface{}{"users:read:*", "users:read:own:other", true},
		[]interface{}{"users:write:*", "users:read:own", false},
		[]interface{}{"users:*:bar", "users:baz:bar", true},
		[]interface{}{"users:*:bar", "users:baz:baz:bar", false},
		[]interface{}{"users:foo:*", "users:foo:bar:foo", true},
	}

	for _, testData := range testingData {
		scopeA := testData.([]interface{})[1].(string)
		scopeB := testData.([]interface{})[0].(string)
		actualOutput := testData.([]interface{})[2].(bool)
		output := MatchScopes(scopeA, scopeB)

		msg := fmt.Sprintf("Failed At Matching Scope %s and %s, result should be %t but MatchScope returned %t", scopeA, scopeB, actualOutput, output)
		assert.Equal(t, actualOutput, output, msg)
	}
}
