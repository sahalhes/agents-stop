package main

import (
	"fmt"
	"os"
	"strings"
)

var version = "dev"

func main() {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "help" || args[0] == "--help" || args[0] == "-h" {
		usage()
		return
	}
	if args[0] == "--version" || args[0] == "version" {
		fmt.Printf("agents %s\n", version)
		return
	}
	if args[0] != "stop" {
		fatal("unknown command %q (try: agents help)", args[0])
	}
	dryRun := false
	for _, arg := range args[1:] {
		if arg == "--dry-run" {
			dryRun = true
		} else {
			fatal("unknown option %q", arg)
		}
	}
	if err := stopAgents(dryRun); err != nil {
		fatal("%v", err)
	}
}

func usage() {
	fmt.Println(`agents - stop AI agent processes

Usage:
  agents stop [--dry-run]
  agents --version

stop scans the whole machine for supported AI coding agents and their child
processes. It forcefully terminates matches. Use --dry-run to preview them.`)
}

func fatal(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "agents: "+format+"\n", args...)
	os.Exit(2)
}

type process struct {
	PID  int    `json:"pid"`
	PPID int    `json:"ppid"`
	Name string `json:"name"`
	Args string `json:"args"`
}

var executableNames = map[string]bool{
	"codex": true, "claude": true, "opencode": true, "goose": true,
	"herd": true, "aider": true, "gemini": true, "cline": true, "kiro": true,
}

var commandMarkers = []string{
	"@openai/codex", "codex app-server", "@anthropic-ai/claude-code", "claude-code",
	"opencode-ai", "goose agent", "herd agent", "aider-chat", "@google/gemini-cli",
	"gemini-cli", "cline", "kiro-cli",
}

func isAgent(p process) bool {
	base := strings.ToLower(p.Name)
	if index := strings.LastIndexAny(base, `/\\`); index >= 0 {
		base = base[index+1:]
	}
	if ext := strings.LastIndexByte(base, '.'); ext >= 0 {
		base = base[:ext]
	}
	if executableNames[base] {
		return true
	}
	if base != "node" && base != "nodejs" && base != "python" && base != "python3" {
		return false
	}
	args := " " + strings.ToLower(p.Args) + " "
	for _, marker := range commandMarkers {
		if strings.Contains(args, marker) {
			return true
		}
	}
	return false
}

func stopAgents(dryRun bool) error {
	all, err := listProcesses()
	if err != nil {
		return fmt.Errorf("scan processes: %w", err)
	}
	selected := make(map[int]process)
	for _, p := range all {
		if isAgent(p) {
			selected[p.PID] = p
		}
	}
	changed := true
	for changed {
		changed = false
		for _, p := range all {
			if _, exists := selected[p.PID]; !exists {
				if _, parentSelected := selected[p.PPID]; parentSelected {
					selected[p.PID] = p
					changed = true
				}
			}
		}
	}
	if len(selected) == 0 {
		fmt.Println("No supported agent processes found.")
		return nil
	}
	order := make([]process, 0, len(selected))
	for _, p := range selected {
		order = append(order, p)
	}
	depth := func(p process) int {
		d, parent := 0, p.PPID
		for d < len(selected) {
			ancestor, ok := selected[parent]
			if !ok {
				break
			}
			d++
			parent = ancestor.PPID
		}
		return d
	}
	for i := 0; i < len(order); i++ {
		for j := i + 1; j < len(order); j++ {
			if depth(order[j]) > depth(order[i]) {
				order[i], order[j] = order[j], order[i]
			}
		}
	}
	for _, p := range order {
		if p.PID == os.Getpid() {
			continue
		}
		if dryRun {
			fmt.Printf("Would stop %s (PID %d)\n", p.Name, p.PID)
			continue
		}
		if err := terminate(p.PID); err != nil {
			fmt.Fprintf(os.Stderr, "Could not stop %s (PID %d): %v\n", p.Name, p.PID, err)
		} else {
			fmt.Printf("Stopped %s (PID %d)\n", p.Name, p.PID)
		}
	}
	return nil
}
