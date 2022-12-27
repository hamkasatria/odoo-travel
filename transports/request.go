package transports

type InsertTravel struct {
	Name    string `json:"Name" validate:"required"`
	Contact string `json:"contact" validate:"required"`
}

type UpdateBook struct {
	Name    string `json:"Name"`
	Contact string `json:"Contact"`
}
