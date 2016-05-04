package electme

import(
    "encoding/json"
)

//Office is the basic data type for Elect-me.
type Office struct {
	Id                  int    `json: "Id"`
	FirstName           string `json: "FirstName"`
	LastName            string `json: "LastName"`
	Position            string `json: "Position"`
	Description         string `json: "Description"`
	Requirments         string `json: "Requirments"`
	Salary              string `json: "Salary"`
	Filing              string `json: "Filing"`
	ResidencyInDistrict string `json: "ResidencyInDistrict"`
	PetitionSignatures  string `json: "PetitionSignatures"`
	MinAge              string `json: "MinAge"`
	VoteCount           string `json: "VoteCount"`
}

var topTen []*Office

//need to deprecate all this...
var data = `
[{"Salary": "95$", "VoteCount": "0", "MinAge": "18", "Requirments": "Live in division", "Description": "Part of election board", "PetitionSignatures": "5", "ResidencyInDistrict": "Registered in division at least 30 days before election", "Position": "INSPECTOR OF ELECTION-29-17", "Filing": "$0"}, 

{"Salary": "100$", "VoteCount": "0", "MinAge": "18", "Requirments": "Live in division", "Description": "Part of election board", "PetitionSignatures": "10", "ResidencyInDistrict": "Registered in division at least 30 days before election", "Position": "JUDGE OF ELECTION-29-17", "Filing": "$0"}, 

{"Salary": "100$", "VoteCount": "1", "MinAge": "18", "Requirments": "Live in division", "Description": "Part of election board", "PetitionSignatures": "10", "ResidencyInDistrict": "Registered in division at least 30 days before election", "Position": "JUDGE OF ELECTION-58-44", "Filing": "$0"}, 

{"Salary": "100$", "VoteCount": "1", "MinAge": "18", "Requirments": "Live in division", "Description": "Part of election board", "PetitionSignatures": "10", "ResidencyInDistrict": "Registered in division at least 30 days before election", "Position": "JUDGE OF ELECTION-21-35", "Filing": "$0"}, 

{"Salary": "95$", "VoteCount": "1", "MinAge": "18", "Requirments": "Live in division", "Description": "Part of election board", "PetitionSignatures": "5", "ResidencyInDistrict": "Registered in division at least 30 days before election", "Position": "INSPECTOR OF ELECTION-59-14", "Filing": "$0"}, 

{"Salary": "95$", "VoteCount": "1", "MinAge": "18", "Requirments": "Live in division", "Description": "Part of election board", "PetitionSignatures": "5", "ResidencyInDistrict": "Registered in division at least 30 days before election", "Position": "INSPECTOR OF ELECTION-2-14", "Filing": "$0"}, 

{"Salary": "95$", "VoteCount": "1", "MinAge": "18", "Requirments": "Live in division", "Description": "Part of election board", "PetitionSignatures": "5", "ResidencyInDistrict": "Registered in division at least 30 days before election", "Position": "INSPECTOR OF ELECTION-41-22", "Filing": "$0"}, 

{"Salary": "100$", "VoteCount": "1", "MinAge": "18", "Requirments": "Live in division", "Description": "Part of election board", "PetitionSignatures": "10", "ResidencyInDistrict": "Registered in division at least 30 days before election", "Position": "JUDGE OF ELECTION-50-27", "Filing": "0$"}, 

{"Salary": "95$", "VoteCount": "1", "MinAge": "18", "Requirments": "Live in division", "Description": "Part of election board", "PetitionSignatures": "5", "ResidencyInDistrict": "Registered in division at least 30 days before election", "Position": "INSPECTOR OF ELECTION-22-29", "Filing": "$0"}, 

{"Salary": "95$", "VoteCount": "1", "MinAge": "18", "Requirments": "Live in division", "Description": "Part of election board", "PetitionSignatures": "5", "ResidencyInDistrict": "Registered in division at least 30 days before election", "Position": "INSPECTOR OF ELECTION-17-16", "Filing": "$0"}]
`

func getOffices() {
	tmp := []byte(data)

	if err := json.Unmarshal(tmp, &topTen); err != nil {
		panic(err)
	}

	//for _,val:= range
}
