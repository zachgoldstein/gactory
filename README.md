# Gactory

A go library for generating new objects will fields populated according to field tags.

## Demo

```golang

default, err := gactory.Fill(SimpleStruct{})
//default now contains values in it's fields corresponding to the field tags

cleanDefault, err := gactory.Convert( newObj )
//cleanDefault copies default into another type that does not have any field tags
//use this when you don't want your structs polluted with field tags  

type SimpleStruct struct {
	Link		     CleanSimpleStruct  `gactory:"GactoryLink"`
	Stuff            string             `gactory:"more stuff happened"`
	Things           int                `gactory:"20"`
	BigOlIntSlice    []int              `gactory:"1,2,3,4"`
}

type CleanSimpleStruct struct {
	Stuff            string
	Things           int
	BigOlIntSlice    []int
}
```

## Rationale

When writing tests, sometimes you write simple functions that just return a struct with certain fields set to testing values.
 
One alternative is to put those values on the field tags instead, and use reflection to generate the values. This library facilitates that.

Since this library heavily uses reflection, it's probably not advisable to use this beyond testing in production.

## Install

` go get github.com/zachgoldstein/gactory `

## Test

` go test `