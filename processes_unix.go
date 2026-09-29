//go:build !windows

package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func listProcesses() ([]process, error) {
	output, err := exec.Command("ps", "-axo", "pid=,ppid=,comm=,args=").Output()
	if err != nil {
		return nil, err
	}
	var result []process
	for _, line := range strings.Split(string(output), "\n") {
		line = strings.TrimSpace(line)
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}
		pid, e1 := strconv.Atoi(fields[0])
		ppid, e2 := strconv.Atoi(fields[1])
		if e1 != nil || e2 != nil {
			continue
		}
		argsStart := len(fields[0]) + len(fields[1])
		for argsStart < len(line) && line[argsStart] == ' ' {
			argsStart++
		}
		args := strings.TrimSpace(line[argsStart:])
		result = append(result, process{PID: pid, PPID: ppid, Name: fields[2], Args: args})
	}
	return result, nil
}

func terminate(pid int) error {
	if err := syscall.Kill(pid, syscall.SIGKILL); err != nil {
		return fmt.Errorf("kill: %w", err)
	}
	return nil
}
