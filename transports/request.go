package transports

type InsertTravel struct {
	Name    string `json:"name" validate:"required"`
	Contact string `json:"contact" validate:"required"`
}

type UpdateTravel struct {
	Name    string `json:"name"`
	Contact string `json:"contact"`
}
