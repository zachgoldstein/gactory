package gactory

import (
	"testing"
	c "github.com/smartystreets/goconvey/convey"
	"log"
	"encoding/json"
)

func TestGenMock(t *testing.T) {
	compare := NewDummyStruct()
	c.Convey("With a struct that has valid field tags", t, func() {
		c.Convey("I can get an instance with the values I expect", func() {
			mockDummyRaw, errs := GenMock(DummyStruct{})
			c.So(len(errs), c.ShouldEqual, 0)

			json, _ := json.MarshalIndent(mockDummyRaw, "", "	")

			log.Printf("mockDummyRaw %v",string(json) )
			dummyObj, ok := mockDummyRaw.(DummyStruct)
			c.So(ok, c.ShouldBeTrue)
			c.So(compare, c.ShouldResemble, dummyObj)
		})
	})
	c.Convey("With a struct that does not have valid field tags", t, func() {
		c.Convey("If the field tags are missing or badly named, the struct does not generate them", func() {
			mockMalformedDummy, errs := GenMock(MalformedStruct{})
			c.So(len(errs), c.ShouldEqual, 0)
			json, _ := json.MarshalIndent(mockMalformedDummy, "", "	")
			log.Printf("mockMalformedDummy %v",string(json) )
		})
		//		c.Convey("If the field tags have non-settable fields, the struct does not generate them", func() {
		//		})
		//		c.Convey("If the field tags have multiple issues, they are all returned", func() {
		//		})
	})
}

func TestMakeValue(t *testing.T) {

}

type MalformedStruct struct {
	Stuff           string  `factori:"stuff happened"`
	things          int     `factory:"10"`
	AwesomeSauce    float64 ``
}

type DummyStruct struct {
	Things          int     `gactory:"10"`
	Stuff           string  `gactory:"stuff happened"`
	AwesomeSauce    float64 `gactory:"1.52"`
	SubStruct		SubObj
}

type SubObj struct {
	Stuff           string  `gactory:"more stuff happened"`
	Things          int     `gactory:"20"`
	AwesomeSauce    float64 `gactory:"2.52"`
}

func NewDummyStruct() DummyStruct {
	return DummyStruct{
		Stuff: "stuff happened",
		Things: 10,
		AwesomeSauce: 1.52,
		SubStruct: SubObj{
			Stuff : "more stuff happened",
			Things: 20,
			AwesomeSauce: 2.52,
		},
	}
}
