package main

type Entity struct {
	Uid     string `json:"uid"`
	Surname string `json:"surname"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}

type Student struct {
	Entity
	Groups []string `json:"groups"`
}

type Info map[string]string
type Roles map[string]string

type Group struct {
	Uid    string `json:"uid"`
	Name   string `json:"name"`
	Preset string `json:"preset"`
	Roles  `json:"roles"`
	Info   `json:"info"`
}

func init() {
	student1 := Student{
		Entity: Entity{
			Uid:     "Aj8BSeViS68p",
			Surname: "Jones",
			Name:    "David-Robert",
			Email:   "david@davidbowie.com",
		},
		Groups: []string{"xW1AkswBapQU"},
	}

	group1 := Group{}
}

func main() {

}
