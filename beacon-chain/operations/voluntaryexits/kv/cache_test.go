package kv

import (
	"testing"

	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
)

func TestKV_SaveGetDelete(t *testing.T) {
	cache := NewCache()
	ve1 := &ethpb.SignedVoluntaryExit{Exit: &ethpb.VoluntaryExit{Epoch: 0, ValidatorIndex: 0}}
	ve2 := &ethpb.SignedVoluntaryExit{Exit: &ethpb.VoluntaryExit{Epoch: 0, ValidatorIndex: 1}}
	ve3 := &ethpb.SignedVoluntaryExit{Exit: &ethpb.VoluntaryExit{Epoch: 0, ValidatorIndex: 2}}

	cacheVes := cache.VoluntaryExits()
	if len(cacheVes) != 0 {
		t.Fatalf("Expected 0 voluntary exits, found %v", len(cacheVes))
	}

	if err := cache.SaveVoluntaryExit(ve1); err != nil {
		t.Fatal(err)
	}
	cacheVes = cache.VoluntaryExits()
	if len(cacheVes) != 1 {
		t.Fatalf("Expected 1 voluntary exit, found %v", len(cacheVes))
	}

	if err := cache.SaveVoluntaryExit(ve2); err != nil {
		t.Fatal(err)
	}
	cacheVes = cache.VoluntaryExits()
	if len(cacheVes) != 2 {
		t.Fatalf("Expected 2 voluntary exits, found %v", len(cacheVes))
	}

	if err := cache.SaveVoluntaryExit(ve3); err != nil {
		t.Fatal(err)
	}
	cacheVes = cache.VoluntaryExits()
	if len(cacheVes) != 3 {
		t.Fatalf("Expected 3 voluntary exits, found %v", len(cacheVes))
	}

	if err := cache.DeleteVoluntaryExit(ve2); err != nil {
		t.Fatal(err)
	}
	cacheVes = cache.VoluntaryExits()
	if len(cacheVes) != 2 {
		t.Fatalf("Expected 2 voluntary exits, found %v", len(cacheVes))
	}
}
