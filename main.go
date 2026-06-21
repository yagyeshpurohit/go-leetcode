package main

import (
	"encoding/json"
	"fmt"
)

type Location struct {
	LocationId int64 `json:"locationId"`
}

type RolesLocation struct {
	Role     string     `json:"role"`
	SubRole  string     `json:"subrole"`
	Location []Location `json:"location"`
}

type ImmediateReporteesDetails struct {
	Id                  int64                  `json:"id"`
	FirstName           string                 `json:"first_name"`
	LastName            string                 `json:"last_name"`
	PrimaryMobileNumber string                 `json:"primary_mobile_number"`
	Metadata            map[string]interface{} `json:"meta_data"`
	Roles               []RolesLocation        `json:"roles"`
}

func main() {
	//var records []*ImmediateReporteesDetails
	var metadata map[string]string
	strJSON := "{\"ownership\": \"INTERNAL\", \"maritalstatus\": \"SINGLE\", \"emergencyDetails\": [{\"name\": \"cont\", \"relation\": \"Father\", \"mobileNumber\": \"3012202531\"}], \"highestQualification\": \"BA\"}"
	err := json.Unmarshal([]byte(strJSON), &metadata)
	if err != nil {
		fmt.Println(err.Error())
	}
	//metadataMap := make(map[string]string)
	//if record.Metadata != "" {
	//	err = json.Unmarshal([]byte(record.Metadata), &metadataMap)
	//}
	fmt.Println(metadata["emergencyDetails"])
	//fmt.Println(records)
}
