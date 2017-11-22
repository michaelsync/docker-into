package main

import(
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSomething(t *testing.T) {

	Convey("Comparing two variables", t, func() {
		myVar := "Hello, world!"
	
		Convey(`"Asdf" should NOT equal "qwerty"`, func() {
			So("Asdf", ShouldNotEqual, "qwerty")
		})
	
		Convey("myVar should not be nil", func() {
			So(myVar, ShouldNotBeNil)
		})

		Convey("This test should be failed", func() {
			So(myVar, ShouldBeNil)
		})
	})
	
}