package complex

import (
	"fmt"
	"testing"
	"time"

	"github.com/prashantv/gostub"
	. "github.com/smartystreets/goconvey/convey"
)

var timeNow = time.Now

var counter = 100
var CleanUp = cleanUp

func cleanUp(val string) {
	fmt.Println(val)
}

func TestFuncDemo(t *testing.T) {
	Convey("TestFuncDemo", t, func() {
		Convey("for succ", func() {
			stubs := gostub.Stub(&counter, 200)
			defer stubs.Reset()
			stubs.Stub(&timeNow, func() time.Time {
				return time.Date(2015, 6, 1, 0, 0, 0, 0, time.UTC)
			})
			stubs.StubFunc(&CleanUp)
			fmt.Println(counter)
			fmt.Println(timeNow().Day())
			CleanUp("Hello go")
		})
	})
}
