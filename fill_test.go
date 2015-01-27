package gactory

import (
	c "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFill(t *testing.T) {
	expect := NewTestStruct()
	c.Convey("With a struct that has valid field tags", t, func() {
		testStruct := TestStruct{}

		c.Convey("I can generate an instance with the values I expect", func() {
			testStruct, err := Fill(testStruct)
			c.So(err, c.ShouldBeNil)

			testAssertStruct, ok := testStruct.(TestStruct)
			c.So(ok, c.ShouldBeTrue)
			c.So(expect, c.ShouldResemble, testAssertStruct)
		})
	})

	c.Convey("With a struct that does not have valid field tags", t, func() {
		_, err := Fill(MalformedStruct{})

		c.Convey("An error occurs if there's errors when setting the field", func() {
			c.So(err, c.ShouldNotBeNil)
		})
	})
}

type MalformedStruct struct {
	Stuff        string  `gactory:"stuff happened"`
	things       int     `gactory:"10f21t"`
	AwesomeSauce float64 ``
}

type TestStruct struct {
	Things           int       `gactory:"10"`
	Stuff            string    `gactory:"stuff happened"`
	AwesomeSauce     float64   `gactory:"1.52"`
	BigOlStrSlice    []string  `gactory:"test,testing,rad stuff"`
	BigOlIntSlice    []int     `gactory:"1,2,3,4"`
	BigOlFloatSlice []float64 `gactory:"1.52,1.53,1.999999"`
	SubStruct        SimpleStruct
}

func NewTestStruct() TestStruct {
	return TestStruct{
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
