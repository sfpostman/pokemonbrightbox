package books

import (
	"encoding/json"
	"example.com/e2e-library/sdk/internal/unmarshal"
	"example.com/e2e-library/sdk/param"
)

type CheckoutNewBookRequest struct {
	CheckedOut *param.Nullable[bool] `json:"checkedOut,omitempty" xml:"checkedOut,omitempty"`
}

func (c CheckoutNewBookRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CheckoutNewBookRequest to string"
	}
	return string(jsonData)
}

func (c *CheckoutNewBookRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}
