package gactory

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGenMock(t *testing.T) {
	expect := NewDummyStruct()
	c.Convey("With a struct that has valid field tags", t, func() {
		testStruct := DummyStruct{}

		c.Convey("I can generate an instance with the values I expect", func() {
			testStruct, err := Make(testStruct)
			c.So(err, c.ShouldBeNil)

			testAssertStruct, ok := testStruct.(DummyStruct)
			c.So(ok, c.ShouldBeTrue)
			c.So(expect, c.ShouldResemble, testAssertStruct)
		})
	})

	c.Convey("With a struct that does not have valid field tags", t, func() {
		_, err := Make(MalformedStruct{})

		c.Convey("An error occurs if there's errors when setting the field", func() {
			c.So(err, c.ShouldNotBeNil)
		})
	})

	c.Convey("With a struct that is linked to another type", t, func() {

		linkTest := NewTestCleanSimpleStruct()
		testStruct := SimpleStruct{}

		c.Convey("I can generate an instance of the linked type with the values I expect", func() {
			newObj, err := Make(testStruct)
			c.So(err, c.ShouldBeNil)

			testStruct, err := Convert( newObj )
			c.So(err, c.ShouldBeNil)

			testAssertStruct, ok := testStruct.(CleanSimpleStruct)
			c.So(ok, c.ShouldBeTrue)
			c.So(testAssertStruct, c.ShouldResemble, linkTest)
		})

		c.Convey("I cannot generate an instance of the wrong linked type", func() {
			newObj, err := Make(testStruct)
			c.So(err, c.ShouldBeNil)

			testStruct, err := Convert( newObj )
			c.So(err, c.ShouldBeNil)

			testAssertStruct, ok := testStruct.(CleanSimpleStruct)

			expect := NewDummyStruct()

			c.So(ok, c.ShouldBeTrue)
			c.So(testAssertStruct, c.ShouldNotResemble, expect)
		})
	})
}

type MalformedStruct struct {
	Stuff        string  `gactory:"stuff happened"`
	things       int     `gactory:"10f21t"`
	AwesomeSauce float64 ``
}

type DummyStruct struct {
	Things           int       `gactory:"10"`
	Stuff            string    `gactory:"stuff happened"`
	AwesomeSauce     float64   `gactory:"1.52"`
	BigOlStrSlice    []string  `gactory:"test,testing,rad stuff"`
	BigOlIntSlice    []int     `gactory:"1,2,3,4"`
	BigOlFloatSlice []float64 `gactory:"1.52,1.53,1.999999"`
	SubStruct        SimpleStruct
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

func NewDummyStruct() DummyStruct {
	return DummyStruct{
		Stuff:        "stuff happened",
		Things:       10,
		AwesomeSauce: 1.52,
		BigOlStrSlice: []string{"test","testing","rad stuff"},
		BigOlIntSlice: []int{1,2,3,4},
		BigOlFloatSlice: []float64{1.52,1.53,1.999999},
		SubStruct: SimpleStruct{
			Stuff:        "more stuff happened",
			Things:       20,
			AwesomeSauce: 2.52,
		},
	}
}
