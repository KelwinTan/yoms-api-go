package domain

type Event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type AllEvents []Event

var Events = AllEvents{
	{
		ID:          "b184bb6e-ac85-4ccb-aae5-6083762d883d",
		Title:       "Intro to Golang",
		Description: "Learn more",
	},
}
