package books

import (
	"encoding/json"
	"example.com/e2e-library/sdk/internal/unmarshal"
	"example.com/e2e-library/sdk/param"
)

type CreateANewBookRequest struct {
	Title         *param.Nullable[string] `json:"title,omitempty" xml:"title,omitempty"`
	Author        *param.Nullable[string] `json:"author,omitempty" xml:"author,omitempty"`
	Genre         *param.Nullable[string] `json:"genre,omitempty" xml:"genre,omitempty"`
	YearPublished *param.Nullable[string] `json:"yearPublished,omitempty" xml:"yearPublished,omitempty"`
}

func (c CreateANewBookRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateANewBookRequest to string"
	}
	return string(jsonData)
}

func (c *CreateANewBookRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}
