package application

import (
	"log"
	"protobuf-go/protogen/application/job"
	"protobuf-go/protogen/application/jobsearch"
	"protobuf-go/protogen/application/software"

	"google.golang.org/protobuf/encoding/protojson"
)

func JobApplication() {
	candidate := jobsearch.JobCandidate{
		JobCandidateId: 1,
		Application: &job.Application{
			Name:     "Kenzo",
			Position: "Manager",
		},
	}

	jCandidateBytes, _ := protojson.Marshal(&candidate)
	jCandidateString := string(jCandidateBytes)

	log.Println(jCandidateString)

}

func SoftwareApplication() {
	software := jobsearch.JobSoftware{
		SoftwareVersionId: 1.44,
		Application: &software.Application{
			Name:      "Ghost Hunter",
			Download:  267834,
			Platforms: []string{"iOS", "Android"},
		},
	}

	jSoftwareBytes, _ := protojson.Marshal(&software)
	jSoftwareString := string(jSoftwareBytes)
	log.Println(jSoftwareString)
}
