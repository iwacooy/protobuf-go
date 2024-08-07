package basic

import (
	"log"
	"protobuf-go/protogen/basic"
)

func BasicUser() {
	scalar := basic.User{
		Id:       1,
		Username: "Ewong",
		IsActive: true,
		Password: []byte("anjayani"),
	}
	log.Println(&scalar)

}
