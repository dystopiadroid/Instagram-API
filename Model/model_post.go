package Model

type Post struct {
	Id string `json:"id,omitempty"`

	Caption string `json:"caption,omitempty"`

	ImgUrl string `json:"imgUrl,omitempty"`

	Timestamp string `json:"timestamp,omitempty"`
}
