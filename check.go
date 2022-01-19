package identitycheck

import (
	"encoding/json"
	"fmt"
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
func (i *IdentityCheck) Check(name string) (*[]Person, error) {
	// create metaphone hash for given name
	hash := phonetics.EncodeMetaphone(name)

	var ps []Person

	f := func(k []byte) error {
		kp := strings.Split(string(k), ".")
		ik := []byte("data." + kp[2])

		// read person data from database
		po, err := i.person.Get(ik)

		if err != nil {
			if err == bitcask.ErrKeyNotFound {
				return fmt.Errorf("the key %s does not exist", string(ik))
			}

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
		return nil, err
	}

	return &ps, nil
}
