package identitycheck

import "errors"

// Country is person country detail
type Country struct {
	Code string `json:"code"`
	Iso  string `json:"iso"`
	Name string `json:"name"`
}

// NewCountry creates country by given name
func NewCountry(name string) (*Country, error) {
	if len(name) == 0 {
		return nil, errors.New("country name is required")
	}

	c := new(Country)

	return c, nil
}
