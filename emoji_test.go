/*
 * Bamboo - A Vietnamese Input method editor
 * Copyright (C) 2018 Luong Thanh Lam <ltlam93@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */
package main

import (
	"testing"
	"unicode/utf8"

	"ibus-bamboo/config"

	ibus "github.com/BambooEngine/goibus"
)

func TestEmojiFindResult(t *testing.T) {
	if trie, err := loadEmojiOne(DictEmojiOne); err == nil {
		emojiTrie.Store(trie)
	}
	var be = NewEmojiEngine()
	if be.MatchString(":'") != true {
		t.Errorf("Finding result for emoji :', expected true, got %v", be.MatchString(":'"))
	}
	if be.MatchString(":')") != true {
		t.Errorf("Finding result for emoji :'), expected true, got %v", be.MatchString(":')"))
	}
	if be.MatchString("gri") != true {
		t.Errorf("Finding result for emoji gri, expected true, got %v", be.MatchString("gri"))
	}
	if be.MatchString("grin") != true {
		t.Errorf("Finding result for emoji grin, expected true, got %v", be.MatchString("grinning"))
	}
}

func TestFilterEmoji(t *testing.T) {
	if trie, err := loadEmojiOne(DictEmojiOne); err == nil {
		emojiTrie.Store(trie)
	}
	var be = NewEmojiEngine()
	var grinnings = be.Filter(":')")
	if !inStringList(grinnings, "😂") {
		t.Errorf("Filtering emojo :'), expected %v, got %v", true, inStringList(grinnings, "😂"))
	}
	var grinnings2 = be.Filter(":")
	if !inStringList(grinnings2, "😂") {
		t.Errorf("Filtering emojo :, expected %v, got %v", true, inStringList(grinnings2, "😂"))
	}
	var grinnings3 = be.Filter("grin")
	if !inStringList(grinnings3, "😀") {
		t.Errorf("Filtering emojo `grin`, expected %v got %v", true, inStringList(grinnings3, "😀"))
	}
}

// The preedit cursor position is measured in characters; using the byte
// length of a multi-byte emoji pushes it past the end of the preedit text.
func TestEmojiPreeditCursorPosCountsRunes(t *testing.T) {
	trie, err := loadEmojiOne(DictEmojiOne)
	if err != nil {
		t.Skipf("emoji data unavailable: %s", err)
	}
	emojiTrie.Store(trie)
	e, fe := newRegressionEngine(config.PreeditIM)
	e.emojiLookupTable = ibus.NewLookupTable()
	e.emoji.keys = []rune("grin")
	e.updateEmojiLookupTable()
	text := fe.getPreeditText()
	if text == "" {
		t.Fatalf("setup: no emoji preedit produced")
	}
	if got, want := fe.getPreeditCursorPos(), uint32(utf8.RuneCountInString(text)); got != want {
		t.Errorf("preedit cursor_pos = %d (%d bytes), want %d runes for %q", got, len(text), want, text)
	}
}

// ProcessKeyEvent already holds the engine lock when it dispatches to the
// emoji handler, so arrow navigation must not call the locking D-Bus methods.
func TestEmojiLookupTableArrowKeysDoNotDeadlock(t *testing.T) {
	e, _ := newRegressionEngine(config.PreeditIM)
	e.isEmojiLTOpened = true
	lt := ibus.NewLookupTable()
	lt.PageSize = uint32(EmojiMaxPageSize)
	for i := 0; i < 5; i++ {
		lt.AppendCandidate("x")
	}
	e.emojiLookupTable = lt

	for _, key := range []uint32{IBusLeft, IBusRight, IBusUp, IBusDown, IBusPageUp, IBusPageDown} {
		runWithDeadlockGuard(t, func() {
			if ret, _ := e.ProcessKeyEvent(key, 0, 0); !ret {
				t.Errorf("key 0x%04x should be consumed by the open emoji table", key)
			}
		})
	}
}
