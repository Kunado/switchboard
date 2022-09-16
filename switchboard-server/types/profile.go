package types

type Profile struct {
	Id      int
	Name    string
	Enabled bool
}

type ProfileName struct {
	Name string `json:"name"`
}
