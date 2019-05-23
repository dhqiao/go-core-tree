package computer

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)


func TestAdd3(t *testing.T) {
	Convey("go test from computer_a_test.go", t, func() {
		So(Add(1, 2), ShouldEqual, 3)
	})
}

