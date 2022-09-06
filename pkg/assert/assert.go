package assert

import (
	"fmt"
	"reflect"
	"testing"
)

func NoError(t *testing.T, err error, msgAndArgs ...interface{}) {
	if err == nil {
		return
	}
	t.Helper()
	fatalf(t, "err: %v%s", err, format(msgAndArgs...))
}

func Equal(t *testing.T, want, got any, msgAndArgs ...interface{}) {
	if equal(want, got) {
		return
	}
	t.Helper()
	fatalf(t, "want != got\n%#v\n%#v%s", want, got, format(msgAndArgs...))
}

func equal(a, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if reflect.DeepEqual(a, b) {
		return true
	}
	aValue := reflect.ValueOf(a)
	bValue := reflect.ValueOf(b)
	return aValue == bValue
}

func format(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 {
		return ""
	}
	return fmt.Sprintf("\n"+msgAndArgs[0].(string), msgAndArgs[1:]...)
}
