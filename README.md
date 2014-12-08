# gactory # 

a go library for quickly grabbing mocks of structs

## install ##

` go get github.com/zachgoldstein/gactory `

## example ##

```golang

type dummystruct struct {
    stuff           string  `factory:"stuff happened"`
    things          int     `factory:"10"`
    awesomesauce    float64 `factory:"1.52"`
}

dummyobject := gactory.get(dummystruct)

```
