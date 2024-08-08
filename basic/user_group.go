package basic

import (
	"log"
	"protobuf-go/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func UserGroup() {
	fanny := basic.User{
		Id:       1,
		Username: "Fanny",
		IsActive: true,
		Password: []byte("bapakkaugarena"),
		Gender:   basic.Gender_GENDER_FEMALE,
		Emails:   []string{"fanny@example.com"},
		Address: &basic.Address{
			Street:  "Jl.Mawar",
			City:    "Tulungagung",
			Country: "Jawa Timur",
		},
	}

	gatot := basic.User{
		Id:       2,
		Username: "Gatot",
		IsActive: true,
		Password: []byte("anjaymabar"),
		Gender:   basic.Gender_GENDER_MALE,
		Emails:   []string{"gatot@example.com"},
		Address: &basic.Address{
			Street:  "Jl.Melati",
			City:    "Sidoarjo",
			Country: "Jawa Timur",
		},
	}

	hero := basic.UserGroup{
		GroupId:     1,
		GroupName:   "Nusantara",
		Roles:       []string{"Admin", "Manager"},
		Users:       []*basic.User{&fanny, &gatot},
		Description: "Hero Meler",
	}

	jsonBytes, _ := protojson.Marshal(&hero)
	jsonString := string(jsonBytes)

	log.Println(jsonString)
}
