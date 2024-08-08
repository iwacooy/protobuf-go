package basic

import (
	"log"
	"protobuf-go/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUser() {
	u := basic.User{
		Id:       1,
		Username: "chilli",
		IsActive: true,
		Password: []byte("anjayani"),
		Gender:   basic.Gender_GENDER_FEMALE,
		Emails:   []string{"chilli@example.com"},
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
