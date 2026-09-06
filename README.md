# ibus-bamboo — correctness audit fork

Fork of [BambooEngine/ibus-bamboo](https://github.com/BambooEngine/ibus-bamboo).
It carries a small correctness audit of the engine: focus handling, the key
queue, startup, session-bus use, and X11 memory safety. Each fix was
reproduced with a failing test before it was made.

The work is up for review upstream as a stack, in merge order:

1. [Settle composition state on focus change, reset, disable](https://github.com/BambooEngine/ibus-bamboo/pull/611)
2. [Give each engine its own key queue; serialize engine state](https://github.com/BambooEngine/ibus-bamboo/pull/612)
3. [Load lookup data without racing or crashing; finish engine setup in the constructor](https://github.com/BambooEngine/ibus-bamboo/pull/613)
4. [Stop closing the shared session-bus connection per focus event](https://github.com/BambooEngine/ibus-bamboo/pull/614)
5. [Fix X11 NULL-deref, leaks, and unbounded copies](https://github.com/BambooEngine/ibus-bamboo/pull/615)

Branches `audit/*` hold the same stack on this fork. `go vet` is clean;
`go test` and `go test -race` pass (needs Go plus X11/GTK dev headers).
