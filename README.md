# wecycle

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

**Cycle your Wallpaper Engine Workshop wallpapers with Alt+Shift on Linux.**

A tiny hotkey companion for [linux-wallpaperengine](https://github.com/Almamu/linux-wallpaperengine). Scans your Steam Workshop folder, keeps track of history, and tells `linux-wallpaperengine` to apply the next wallpaper with a single keybind.

## Why

[linux-wallpaperengine](https://github.com/Almamu/linux-wallpaperengine) does a great job of rendering Wallpaper Engine's Workshop wallpapers on Linux, but there's no way to just press `Alt+Shift` to go to the next one. `wecycle` fixes that.

## Features

- **Hotkey friendly** - Bind `Alt+Shift` to `wecycle next` and you're done
- **Uses your existing workshop** - Scans `steamapps/workshop/content/431960/` automatically
- **Random or sequential** - `random` by default, or cycle in order
- **History aware** - `wecycle prev` works
- **Works with linux-wallpaperengine** - Calls `linux-wallpaperengine --directory <project>` to apply
- **Multi Steam path support** - Checks `~/.steam`, `~/.local/share/Steam`, and custom locations
- **Minimal & fast** - Single binary, no GUI

## Requirements

- [Wallpaper Engine](https://store.steampowered.com/app/431960/Wallpaper_Engine/) installed via Steam with wallpapers subscribed
- [linux-wallpaperengine](https://github.com/Almamu/linux-wallpaperengine) installed and working

## Install

```bash
go install github.com/UnsoberDriver/wecycle/cmd/wecycle@latest
```

Or build from source:

```bash
git clone https://github.com/UnsoberDriver/wecycle
cd wecycle
go build ./cmd/wecycle
```

## Quick start

```bash
# 1) Create config (optional, auto-detects Steam paths)
mkdir -p ~/.config/wecycle
wecycle init > ~/.config/wecycle/config.toml

# 2) Test
wecycle list    # see detected workshop wallpapers
wecycle next    # apply next wallpaper
wecycle random  # apply random wallpaper

# 3) Bind a hotkey
# KDE Plasma: System Settings > Shortcuts > Custom Shortcuts > Add Command/URL
#   Command: /home/youruser/go/bin/wecycle next
#   Trigger: Alt+Shift

# Hyprland: bind = ALT SHIFT, W, exec, wecycle next
# Sway: bindsym Mod1+Shift+w exec wecycle next
```

## Commands

| Command | Description |
| --- | --- |
| `wecycle list` | List all detected Wallpaper Engine workshop wallpapers |
| `wecycle next` | Go to next wallpaper (respects mode from config) |
| `wecycle random` | Pick and apply a random wallpaper |
| `wecycle prev` | Go back to previous wallpaper |
| `wecycle init` | Print a starter config to stdout |

## Config

```toml
[settings]
mode = "random"        # "random" or "sequential"
workshop_dir = ""      # leave empty for auto-detect
exclude = []           # workshop IDs to skip

[linux_wallpaperengine]
binary = "linux-wallpaperengine"  # path to binary if not in PATH
args = []                         # extra args to pass
```

## Tips

- If `wecycle list` returns nothing, make sure you have wallpapers subscribed in Steam and they've been downloaded.
- `wecycle` just picks the project folder and tells `linux-wallpaperengine` to run it. It doesn't download anything from the Workshop.
- Works great alongside `linux-wallpaperengine` running as your wallpaper renderer.

## License

MIT
