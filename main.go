package main

import dtoAlias "example.com/gowrap/issue/98/dto"

//go:generate go run github.com/hexdigest/gowrap/cmd/gowrap gen -i MyInterface -t prometheus -o prometheus_gen.go -g
type MyInterface interface {
	SomeMethod(dto dtoAlias.MyDTO) error
}

func main() {

}
