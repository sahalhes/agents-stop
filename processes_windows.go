//go:build windows

package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

func listProcesses() ([]process, error) {
	script := `Get-CimInstance Win32_Process | Select-Object @{n='pid';e={$_.ProcessId}},@{n='ppid';e={$_.ParentProcessId}},@{n='name';e={$_.Name}},@{n='args';e={$_.CommandLine}} | ConvertTo-Json -Compress`
	output, err := exec.Command("powershell.exe", "-NoLogo", "-NoProfile", "-NonInteractive", "-Command", script).Output()
	if err != nil {
		return nil, err
	}
	var result []process
	if err := json.Unmarshal(output, &result); err == nil && len(result) > 0 {
		return result, nil
	}
	trimmed := strings.TrimSpace(string(output))
	if trimmed == "" || trimmed == "null" {
		return nil, nil
	}
	var one process
	if err := json.Unmarshal(output, &one); err != nil {
		return nil, fmt.Errorf("decode process list: %w", err)
	}
	return []process{one}, nil
}

func terminate(pid int) error {
	output, err := exec.Command("taskkill.exe", "/PID", fmt.Sprint(pid), "/F").CombinedOutput()
	if err != nil {
		return fmt.Errorf("taskkill: %s", strings.TrimSpace(string(output)))
	}
	return nil
}
