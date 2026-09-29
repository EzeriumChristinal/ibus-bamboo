package main

import (
	"strings"
	"testing"

	"ibus-bamboo/config"

	ibus "github.com/BambooEngine/goibus"
)

// Map iteration order must not leak into the D-Bus property list.
func TestInputMethodMenuOrderIsStable(t *testing.T) {
	cfg := config.DefaultCfg()
	if len(cfg.InputMethodDefinitions) < 2 {
		t.Skip("config defines fewer than two input methods")
	}
	var first string
	for i := 0; i < 30; i++ {
		var keys []string
		for _, v := range GetIMPropListByConfig(&cfg).PropertyList {
			if p, ok := v.Value().(ibus.Property); ok {
				keys = append(keys, p.Key)
			}
		}
		got := strings.Join(keys, ",")
		if i == 0 {
			first = got
			continue
		}
		if got != first {
			t.Fatalf("input-method menu order is unstable:\n first %s\n  then %s", first, got)
		}
	}
}
