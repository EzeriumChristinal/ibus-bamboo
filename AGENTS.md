# AGENTS.md

Read this before touching the tree. Say less, know more: short replies, no narration of obvious steps.

## Repo

- ibus-bamboo: IBus Vietnamese input engine. Go + cgo (X11/Xtst) + a GTK3 UI package.
- Upstream is BambooEngine/ibus-bamboo; `origin` here is the fork. Never push upstream; open cross-repo PRs from `audit/*` branches.
- Fork `master` carries fork-only files (README.md, this file). Feature work goes on stacked `audit/*` branches, one topic each.
- Worktree state you did not create is intentional until told otherwise. Never `git add -A`; stage explicit paths, and never commit a deletion you did not make.

## Build and test in this sandbox

- No Go toolchain and no X11/GTK headers on PATH; `scripts/build` and `scripts/test` only run once you stage them.
- Two environment traps, both confirmed: the workspace path contains a space (breaks cgo/pkg-config flag parsing), and `/tmp` is wiped between shell calls.
- Setup that works: download the Go tarball plus the -dev packages (`apt-get download` needs no sudo) into a workspace-local cache, extract to a space-free path, and build/test from there. Re-extract per shell call from the cache.
- Point `PKG_CONFIG_SYSROOT_DIR` at the extracted root; resolve dangling dev `.so` symlinks against the host runtime libs or linking fails.
- Gate every change on: `go vet`, `go test`, `go test -race`, a full `go build` (covers the C and GTK UI), `gofmt -l`.
- Delete the staging dir when done. Never commit it.

## Working rules

- Prove before fixing: failing test first, then the smallest change that turns it green.
- One topic per PR. Every non-trivial fix gets regression coverage; `fake_engine.go` is the harness, keep it race-safe.
- Fragile, touch lightly: the preedit/commit state machine, focus handlers, backspace timing sleeps.
- IBus/framework bugs get documented, not worked around in-engine.

## Use the tools you have

- `gh` is authenticated: repos, PRs, forks. Prefer it over hand-rolled hosting steps.
- Skills first: match the task against the skill catalog before improvising. Load `deslopify` before writing any PR or repo text.
- `scripts/build` and `scripts/test` are canonical once the toolchain is staged.
- Subagents for long downloads or independent recon while the main line continues.
