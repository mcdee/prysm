package voluntaryexits

import (
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	"github.com/prysmaticlabs/prysm/beacon-chain/operations/voluntaryexits/kv"
)

// Pool defines the necessary methods to manage voluntary exit pools.
type Pool interface {
	SaveVoluntaryExit(exit *ethpb.SignedVoluntaryExit) error
	HasVoluntaryExit(root [32]byte) bool
	DeleteVoluntaryExit(exit *ethpb.SignedVoluntaryExit) error
	VoluntaryExits() []*ethpb.SignedVoluntaryExit
}

// NewPool initializes a new voluntary exit pool.
func NewPool() *kv.Cache {
	return kv.NewCache()
}
