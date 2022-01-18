package identitycheck

import "time"

// Person entry type
const (
	PEP      = "pep"
	SANCTION = "sanction"
	UNKNOWN  = "unknown"
)

// Person contains name, birth detail, country and entry type
type Person struct {
	Hash       string     `json:"hash"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	BirthDay   *int       `json:"birth_day"`
	BirthMonth *int       `json:"birth_month"`
	BirthYear  *int       `json:"birth_year"`
	Alias      *[]string  `json:"alias"`
	BirthDate  *time.Time `json:"birth_date"`
	Country    *Country   `json:"country"`
}
