package basic

import (
	"log"
	"protobuf-go/protogen/basic"
)

func BasicHello() {
	h := basic.Hello{
		Name: "Jess no tititd",
	}

	log.Println(&h)
}
