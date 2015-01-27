package gactory

import (
	"testing"
	c "github.com/smartystreets/goconvey/convey"
)

func TestConvert(t *testing.T) {

	c.Convey("With a struct that is linked to another type", t, func() {

		linkTest := NewTestCleanSimpleStruct()
		testStruct := SimpleStruct{}

		c.Convey("I can generate an instance of the linked type with the values I expect", func() {
			newObj, err := Fill(testStruct)
			c.So(err, c.ShouldBeNil)

			testStruct, err := Convert( newObj )
			c.So(err, c.ShouldBeNil)

			testAssertStruct, ok := testStruct.(CleanSimpleStruct)
			c.So(ok, c.ShouldBeTrue)
			c.So(testAssertStruct, c.ShouldResemble, linkTest)
		})

		c.Convey("I cannot generate an instance of the wrong linked type", func() {
			newObj, err := Fill(testStruct)
			c.So(err, c.ShouldBeNil)

			testStruct, err := Convert( newObj )
			c.So(err, c.ShouldBeNil)

			testAssertStruct, ok := testStruct.(CleanSimpleStruct)

			expect := NewTestStruct()

			c.So(ok, c.ShouldBeTrue)
			c.So(testAssertStruct, c.ShouldNotResemble, expect)
		})
	})
}

type SimpleStruct struct {
	Link		 CleanSimpleStruct `gactory:"GactoryLink"`
	Stuff        string  `gactory:"more stuff happened"`
	Things       int     `gactory:"20"`
	AwesomeSauce float64 `gactory:"2.52"`
}

type CleanSimpleStruct struct {
	Stuff        string
	Things       int
	AwesomeSauce float64
}

func NewTestCleanSimpleStruct() CleanSimpleStruct {
	return CleanSimpleStruct{
		Stuff: "more stuff happened",
		Things: 20,
		AwesomeSauce: 2.52,
	}
}
