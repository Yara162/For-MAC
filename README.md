# For-MAC

**MAC Address Formatter v1.2**: a small, standalone Windows app that turns MAC addresses
into the standard `AA:BB:CC:DD:EE:FF` format, however they were typed or pasted.

Native port of the original single-file HTML tool, with the same behavior.

## Download

Get `For-MAC.exe` from the [Releases](../../releases) page. It's one portable executable:
no installer, no runtime. Windows 7 or later, 64-bit.

## Usage

1. Paste or type MAC addresses into **Input**, one per line.
2. Click **Do** (or press **Ctrl+Enter**).
3. Click **Copy Output** to copy the results.

Each output line lines up with its input line:

| Input                | Output                       |
|----------------------|------------------------------|
| `001a2b3c4d5e`       | `00:1A:2B:3C:4D:5E`          |
| `00-1A-2B-3C-4D-5E`  | `00:1A:2B:3C:4D:5E`          |
| `001a.2b3c.4d5e`     | `00:1A:2B:3C:4D:5E`          |
| `00 1a 2b 3c 4d 5e`  | `00:1A:2B:3C:4D:5E`          |
| *(blank)*            | *(blank)*                    |
| `12345`              | `ERROR: too few characters`  |
| `001122334455667`    | `ERROR: too many characters` |

Rule: every non-hex character is removed. If exactly 12 hex digits remain, they're uppercased
and grouped in pairs with colons.

## Building

Needs Go 1.21+. The build cross-compiles from Linux or macOS, so you don't need a Windows machine.

```bash
./build.sh
```

The output is `dist/For-MAC.exe`. The icon, manifest and version info come from `winres/` and
are compiled into `rsrc_windows_amd64.syso`. To regenerate it after editing `winres/`, install
[go-winres](https://github.com/tc-hib/go-winres) (`go install github.com/tc-hib/go-winres@latest`).

Layout:

- `main.go`: the Windows GUI ([lxn/walk](https://github.com/lxn/walk))
- `internal/macfmt/`: the formatting logic and its unit tests (cross-platform)
- `assets/`: the logo
- `winres/`: icon, manifest and version resource definitions

---

Made by YZ for Smart Spaces WI LLC. CC BY 4.0
