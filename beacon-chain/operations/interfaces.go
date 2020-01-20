package operations

import (
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
)

// VoluntaryExitsPoolFetcher defines the interface for fetching voluntary exits from the local pool.
type VoluntaryExitsPoolFetcher interface {
	// VoluntaryExits returns all voluntary exits in the pool.
	VoluntaryExits() []*ethpb.SignedVoluntaryExit
}
