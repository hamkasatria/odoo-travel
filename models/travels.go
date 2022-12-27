package models

// Travel struct
type Travel struct {
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string `json:"name" bson:"name"`
	Contact string `json:"contact" bson:"contact"`
}
