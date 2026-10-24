// Package main
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/UnsoberDriver/wecycle/internal/cache"
	"github.com/UnsoberDriver/wecycle/internal/config"
	"github.com/UnsoberDriver/wecycle/internal/runner"
	"github.com/UnsoberDriver/wecycle/internal/workshop"
)

var version = "dev"

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	cmd := os.Args[1]
	switch cmd {
	case "version":
		fmt.Println("wecycle", version)
		return
	case "init":
		printInit()
		return
	case "list":
		runList()
		return
	case "next", "random", "prev":
		runSet(cmd)
		return
	default:
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Println(`wecycle - cycle Wallpaper Engine Workshop wallpapers with Alt+Shift

Usage:
  wecycle list    list detected workshop wallpapers
  wecycle next    go to next (random or sequential)
  wecycle random  pick a random wallpaper
  wecycle prev    go to previous
  wecycle init    print starter config
  wecycle version`)
}

func printInit() {
	fmt.Println(`# ~/.config/wecycle/config.toml
[settings]
mode = "random"        # random or sequential
# workshop_dir = ""    # leave empty to auto-detect Steam paths
exclude = []

[linux_wallpaperengine]
binary = "linux-wallpaperengine"
args = []`)
}

func load() (*config.Config, *cache.Cache, *cache.State, []string, error) {
	cfg, err := config.Load("")
	if err != nil {
		return nil, nil, nil, nil, err
	}
	c := cache.New()
	st, err := c.Load()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	var dirs []string
	if cfg.Settings.WorkshopDir != "" {
		dirs = []string{cfg.Settings.WorkshopDir}
	} else {
		dirs = workshop.FindDirs()
	}
	if len(dirs) == 0 {
		return nil, nil, nil, nil, fmt.Errorf("no steam workshop folder found for app 431960. Set settings.workshop_dir in config")
	}

	var all []string
	excl := map[string]bool{}
	for _, e := range cfg.Settings.Exclude {
		excl[e] = true
	}
	for _, d := range dirs {
		list, err := workshop.Scan(d)
		if err != nil {
			continue
		}
		for _, p := range list {
			id := p[stringsLastSlash(p)+1:]
			if excl[id] {
				continue
			}
			all = append(all, p)
		}
	}
	if len(all) == 0 {
		return nil, nil, nil, nil, fmt.Errorf("no wallpapers found in workshop folders: %v", dirs)
	}
	st.All = all
	return cfg, c, st, all, nil
}

func stringsLastSlash(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			return i
		}
	}
	return -1
}

func runList() {
	_, _, _, list, err := load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	for _, p := range list {
		fmt.Println(p)
	}
}

func runSet(cmd string) {
	cfg, c, st, list, err := load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())

	var target string
	if cmd == "random" || (cmd == "next" && cfg.Settings.Mode == "random") {
		target = list[rand.Intn(len(list))]
		if len(list) > 1 && target == st.Last {
			target = list[rand.Intn(len(list))]
		}
	} else if cmd == "prev" && len(st.History) > 1 {
		h := st.History
		target = h[len(h)-2]
	} else {
		if st.Index < 0 || st.Index >= len(list) {
			st.Index = 0
		}
		st.Index = (st.Index + 1) % len(list)
		target = list[st.Index]
	}

	if err := runner.Apply(cfg, target); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	st.Last = target
	st.History = append(st.History, target)
	if len(st.History) > 100 {
		st.History = st.History[len(st.History)-100:]
	}
	if err := c.Save(st); err != nil {
		fmt.Fprintln(os.Stderr, "warning:", err)
	}
	fmt.Println(target)
}
