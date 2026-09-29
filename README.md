# ibus-bamboo (audit fork)

Correctness and memory-safety fixes for [BambooEngine/ibus-bamboo](https://github.com/BambooEngine/ibus-bamboo), one topic per branch, each with a regression test.

Bản fork sửa các lỗi về đồng bộ trạng thái gõ, race condition, bộ nhớ X11 và kết nối DBus. Mỗi lỗi đều có test tái hiện trước khi sửa.

## Patch stack

Nine branches on top of `master`, stacked in merge order: each branch contains every branch above it.

| # | Branch | Fix | Upstream PR |
|---|--------|-----|-------------|
| 1 | `audit/focus-lifecycle` | Commit or drop the composition buffer on focus change, reset and disable. | [#611](https://github.com/BambooEngine/ibus-bamboo/pull/611) |
| 2 | `audit/engine-queue` | Give every engine instance its own key queue and serialize access to engine state. | [#612](https://github.com/BambooEngine/ibus-bamboo/pull/612) |
| 3 | `audit/load-robustness` | Load the emoji trie and spelling dictionary without races; finish engine setup in the constructor. | [#613](https://github.com/BambooEngine/ibus-bamboo/pull/613) |
| 4 | `audit/dbus-integration` | Keep the shared session-bus connection open across focus events. | [#614](https://github.com/BambooEngine/ibus-bamboo/pull/614) |
| 5 | `audit/x11-memsafety` | Fix NULL dereferences, leaks and unbounded copies in the X11 bindings. | [#615](https://github.com/BambooEngine/ibus-bamboo/pull/615) |
| 6 | `audit/reentrant-lock` | Lookup-table navigation re-locked the engine mutex and hung the engine on arrow keys. | pending |
| 7 | `audit/queue-accounting` | Keep the pending key counter in sync when a batch commit drains the queue. | pending |
| 8 | `audit/emoji-cursor` | Report the emoji preedit cursor in characters instead of UTF-8 bytes. | pending |
| 9 | `audit/prop-order` | Emit the input-method property list in a stable order. | pending |

Tip of the stack: `audit/round-2`.

Branches 6 to 9 come from a second audit pass over branch 5. Branch 6 fixes a deadlock introduced by branch 2; the commit message carries the reproducer.

## Verification

Every commit passes on its own:

```sh
gofmt -l .
go vet ./...
go test ./...
go test -race ./...
go build ./...   # also builds the cgo X11 and GTK3 packages
```

`fake_engine.go` records preedit cursor positions, which is what lets the emoji and lock fixes be tested rather than assumed.

## Build & Test

Requirements, Debian / Ubuntu / LMDE:

```sh
sudo apt install golang libibus-1.0-dev libgtk-3-dev libx11-dev
```

Commands:

```sh
# Run test suite with race detector
go test -race ./...

# Build binary
make build
```

## Upstream

- Source: [BambooEngine/ibus-bamboo](https://github.com/BambooEngine/ibus-bamboo)
- License: GPL-3.0
