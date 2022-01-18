package identitycheck

import (
	"log"
	"testing"
)

func TestSimple(t *testing.T) {
	c, err := NewIdentityCheck()

	if err != nil {
		t.Error(err)
	}

	p := new(Person)
	p.Name = "Barack Obama"

	ps, err := c.Check(p)

	if err != nil {
		t.Error(err)
	}

	if len(*ps) == 0 {
		t.Error("empty list")
	}

	for _, p := range *ps {
		log.Println(p.Name)
	}
}
