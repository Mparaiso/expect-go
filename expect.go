package expect

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

// Expect is a helper used to reduce the boilerplate during tests
func Expect(t *testing.T, got, want interface{}, comments ...string) {
	var comment string
	if want != got {
		if len(comments) > 0 {
			comment = comments[0]

		} else {
			comment = "Expect"
		}
		_, file, line, _ := runtime.Caller(1)
		t.Fatalf(fmt.Sprintf("Expect\r%s:%d:\r\t%s : %s", filepath.Base(file), line, comment, "want '%v' got '%v'."), want, got)
	}
}
