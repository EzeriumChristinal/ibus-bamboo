# ibus-bamboo (audit fork)

A correctness-focused fork of [BambooEngine/ibus-bamboo](https://github.com/BambooEngine/ibus-bamboo).

Bản fork sửa các lỗi về đồng bộ trạng thái, race condition, bộ nhớ X11 và kết nối DBus. Mỗi lỗi đều có unit test reproducing trước khi sửa.

---

## Patches / Các bản sửa lỗi

Changes are structured as 5 commits pending upstream review:

1. **Focus state handling** ([PR #611](https://github.com/BambooEngine/ibus-bamboo/pull/611))
   * Settles composition state on focus change, reset, and disable events.
   * Cố định trạng thái gõ khi chuyển window, reset hoặc tắt bộ gõ.

2. **Isolated key queue** ([PR #612](https://github.com/BambooEngine/ibus-bamboo/pull/612))
   * Dedicated key queue per engine instance; serializes state changes.
   * Mỗi engine dùng một key queue riêng, cô lập và xếp hàng thứ tự xử lý phím.

3. **Safe initialization** ([PR #613](https://github.com/BambooEngine/ibus-bamboo/pull/613))
   * Loads lookup data without races; finishes setup inside constructor.
   * Khởi tạo dữ liệu tra cứu an toàn luồng, chuyển cài đặt engine vào constructor.

4. **Persistent DBus session** ([PR #614](https://github.com/BambooEngine/ibus-bamboo/pull/614))
   * Keeps shared session-bus connection open across focus events.
   * Không đóng/mở lại kết nối DBus dùng chung mỗi khi đổi focus.

5. **X11 memory safety** ([PR #615](https://github.com/BambooEngine/ibus-bamboo/pull/615))
   * Fixes NULL dereferences, memory leaks, and unbounded copies in X11 bindings.
   * Sửa lỗi con trỏ NULL, rò rỉ bộ nhớ và tràn đệm trong code X11.

---

## Status / Trạng thái

* **Branches**: `audit/*` contains individual patches.
* **Verification**: `go vet`, `go test`, and `go test -race` pass cleanly.

---

## Build & Test / Biên dịch và Kiểm thử

### Requirements / Thư viện phụ thuộc

Debian / Ubuntu / LMDE:
```bash
sudo apt install golang libibus-1.0-dev libgtk-3-dev libx11-dev
```

### Commands / Lệnh thực thi

```bash
# Run test suite with race detector
go test -race ./...

# Build binary
make build
```

---

## Upstream

* Source: [BambooEngine/ibus-bamboo](https://github.com/BambooEngine/ibus-bamboo)
* License: GPL-3.0
