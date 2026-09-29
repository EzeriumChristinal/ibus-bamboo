package main

import (
	"testing"
	"time"

	"ibus-bamboo/config"

	ibus "github.com/BambooEngine/goibus"
)

// runWithDeadlockGuard fails the test if fn does not return. A regressed
// engine lock makes ProcessKeyEvent block forever, which no assertion can
// observe otherwise.
func runWithDeadlockGuard(t *testing.T, fn func()) {
	t.Helper()
	done := make(chan struct{})
	go func() {
		defer close(done)
		fn()
	}()
	select {
	case <-done:
	case <-time.After(3 * time.Second):
		t.Fatal("ProcessKeyEvent did not return: re-entrant engine lock")
	}
}

// ProcessKeyEvent already holds the engine lock when it dispatches to the
// lookup-table handlers, so those must not call the locking D-Bus methods.
func TestInputModeLookupTableArrowKeysDoNotDeadlock(t *testing.T) {
	e, _ := newRegressionEngine(config.PreeditIM)
	e.wmClasses = "kitty:kitty"
	e.isInputModeLTOpened = true
	lt := ibus.NewLookupTable()
	lt.PageSize = uint32(len(config.ImLookupTable))
	for i := 0; i < len(config.ImLookupTable); i++ {
		lt.AppendCandidate("c")
	}
	e.inputModeLookupTable = lt

	for _, key := range []uint32{IBusLeft, IBusRight, IBusUp, IBusDown, IBusPageUp, IBusPageDown} {
		runWithDeadlockGuard(t, func() {
			if ret, _ := e.ProcessKeyEvent(key, 0, 0); !ret {
				t.Errorf("key 0x%04x should be consumed by the open lookup table", key)
			}
		})
	}
}
