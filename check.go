package identitycheck

import (
	"encoding/json"
	"errors"
	"path/filepath"
	"strings"

	"git.mills.io/prologic/bitcask"
	"github.com/dotcypress/phonetics"
)

// IdentityCheck provides base checking
type IdentityCheck struct {
	person *bitcask.Bitcask
	index  *bitcask.Bitcask
}

// NewIdentityCheck creates new pep and sanction check
func NewIdentityCheck() (*IdentityCheck, error) {
	c := new(IdentityCheck)

	pt := filepath.Join("source", "person")
	it := filepath.Join("source", "person.idx")

	pd, err := bitcask.Open(pt)

	if err != nil {
		return nil, err
	}

	c.person = pd

	id, err := bitcask.Open(it)

	if err != nil {
		return nil, err
	}

	c.index = id

	return c, nil
}

// Check finds given person from PEP and sanctions lists
func (i *IdentityCheck) Check(person *Person) (*[]Person, error) {
	// create metaphone hash for given name
	hash := phonetics.EncodeMetaphone(person.Name)

	var ps []Person

	f := func(k []byte) error {
		kp := strings.Split(string(k), ".")

		// read person data from database
		po, err := i.person.Get([]byte("data." + kp[2]))

		if err != nil {
			return err
		}

		p := new(Person)

		err = json.Unmarshal(po, p)

		if err != nil {
			return err
		}

		ps = append(ps, *p)

		return nil
	}

	// iterate over records by given hash prefix
	err := i.index.Scan([]byte("idx."+hash), f)

	if err != nil {
		if err == bitcask.ErrKeyNotFound {
			return nil, errors.New("")
		}

		return nil, err
	}

	return &ps, nil
}
