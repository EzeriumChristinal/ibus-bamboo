package main

import (
	"sync/atomic"
	"testing"

	"ibus-bamboo/config"
)

// updatePreviousTextInBatch drains keyPressChan inline; every dequeued item
// must decrement keyQueueLen exactly like keyPressCapturing does.
func TestBatchDrainKeepsQueueCounterInSync(t *testing.T) {
	e, _ := newRegressionEngine(config.PreeditIM)
	e.enqueueKeyPress('a', 'a', 0)
	if got := atomic.LoadInt32(&e.keyQueueLen); got != 1 {
		t.Fatalf("setup: keyQueueLen = %d, want 1", got)
	}
	e.updatePreviousTextInBatch("", "a", false)
	if got := atomic.LoadInt32(&e.keyQueueLen); got != 0 {
		t.Errorf("batch drain leaked a queued key: keyQueueLen = %d, want 0", got)
	}
}
