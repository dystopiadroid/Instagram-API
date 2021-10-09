package Model

type User struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Email string `json:"email,omitempty"`

	Password string `json:"password,omitempty"`
}
