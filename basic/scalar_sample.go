package basic

import (
	"log"
	"protobuf-go/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUser() {
	addr := basic.Address{
		Street:  "Jl.Mawar",
		City:    "Tulungagung",
		Country: "Jawa Timur",
		Coordinate: &basic.Address_Coordinate{
			Latitude:  -40.4263334564634,
			Longitude: 54.457455232656,
		},
	}
	u := basic.User{
		Id:       1,
		Username: "chilli",
		IsActive: true,
		Password: []byte("anjayani"),
		Gender:   basic.Gender_GENDER_FEMALE,
		Emails:   []string{"chilli@example.com"},
		Address:  &addr,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))

}

func ProtoToJson() {
	u := basic.User{
		Id:       1,
		Username: "ryujin",
		IsActive: true,
		Password: []byte("anjayani"),
		Gender:   basic.Gender_GENDER_FEMALE,
		Emails:   []string{"ryujin@example.com", "makeit@example.com"},
	}

	jsonBytes, _ := protojson.Marshal(&u)
	jsonString := string(jsonBytes)
	log.Println(jsonString)

	log.Println("=============================================")

	var p basic.User
	err := protojson.Unmarshal([]byte(jsonString), &p)
	if err != nil {
		log.Fatal("Got error!", err)
	}

	log.Println(&p)

}

func JsonToProto() {
	jsonData := `{
		"id": 1,
		"username": "ryujin",
		"is_active": true,
		"password": "YW5qYXlhbmk=",
		"gender": "GENDER_FEMALE",
		"emails": [
		  "ryujin@example.com",
		  "makeit@example.com"
		]
	}`

	var p basic.User
	err := protojson.Unmarshal([]byte(jsonData), &p)
	if err != nil {
		log.Fatal("Got error!", err)
	}

	log.Println(&p)
}
