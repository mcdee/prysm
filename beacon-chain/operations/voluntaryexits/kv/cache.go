package kv

import (
	"github.com/patrickmn/go-cache"
	"github.com/pkg/errors"
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"github.com/prysmaticlabs/go-ssz"
)

// SaveVoluntaryExit saves a voluntary exit in cache.
func (c *Cache) SaveVoluntaryExit(exit *ethpb.SignedVoluntaryExit) error {
	root, err := ssz.HashTreeRoot(exit)
	if err != nil {
		return errors.Wrap(err, "could not tree hash voluntary exit")
	}

	c.cache.Set(string(root[:]), exit, cache.DefaultExpiration)

	return nil
}

// HasVoluntaryExit returns true if the voluntary exit is known in the pool.
func (c *Cache) HasVoluntaryExit(root [32]byte) bool {
	_, exists := c.cache.Get(string(root[:]))
	return exists
}

// VoluntaryExits returns all the voluntary exits in cache.
func (c *Cache) VoluntaryExits() []*ethpb.SignedVoluntaryExit {
	voluntaryExits := make([]*ethpb.SignedVoluntaryExit, 0, c.cache.ItemCount())
	for s, i := range c.cache.Items() {
		// Type assertion for the worst case. This shouldn't happen.
		exit, ok := i.Object.(*ethpb.SignedVoluntaryExit)
		if !ok {
			c.cache.Delete(s)
		}

		voluntaryExits = append(voluntaryExits, exit)
	}

	return voluntaryExits
}

// DeleteVoluntaryExit deletes the voluntary exit in cache.
func (c *Cache) DeleteVoluntaryExit(exit *ethpb.SignedVoluntaryExit) error {
	root, err := ssz.HashTreeRoot(exit)
	if err != nil {
		return errors.Wrap(err, "could not tree hash voluntary exit")
	}

	c.cache.Delete(string(root[:]))

	return nil
}
