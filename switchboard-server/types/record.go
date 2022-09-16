package types

type Record struct {
	Id        int
	Host      string
	Value     string
	ProfileId int
}

type RecordBuilder struct {
	Host        string `json:"host"`
	Value       string `json:"value"`
	ProfileName string `json:"profile_name"`
}

type RecordValue struct {
	Value string `json:"value"`
}
