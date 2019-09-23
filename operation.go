package scopes

import (
	"reflect"
)

const (
	//READ is used to address read operation
	READ = "read"
	//WRITE is is used to address write operation
	WRITE = "write"
)

//IsOperationAllowed checks if the operation a a field of a struct is allowed
func IsOperationAllowed(
	data interface{},
	fieldName string,
	operation string,
	scopesAllowed []string) bool {

	typeOf := reflect.TypeOf(data)
	field, _ := typeOf.FieldByName(fieldName)
	scopeRequired := field.Tag.Get(operation + "Scope")
	return ScopeInAllowed(scopeRequired, scopesAllowed)
}

//FilterRead is used to filter output to onlly output what client can see based on scopes
func FilterRead(data interface{}, scopesAllowed []string) {
	valueOf := reflect.Indirect(reflect.ValueOf(data))
	typeOf := valueOf.Type()
	fieldNum := valueOf.NumField()

	for i := 0; i < fieldNum; i++ {
		curField := typeOf.Field(i)
		scopeRequired := curField.Tag.Get("readScope")
		if !ScopeInAllowed(scopeRequired, scopesAllowed) {
			field := valueOf.Field(i)
			field.Set(reflect.Zero(field.Type()))
		}
	}
}
