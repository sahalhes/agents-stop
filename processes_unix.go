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
	return parsePSOutput(string(output)), nil
}

func parsePSOutput(output string) []process {
	var result []process
	for _, line := range strings.Split(output, "\n") {
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
		args := ""
		start := 0
		for field := 0; field < 3; field++ {
			for start < len(line) && line[start] != ' ' && line[start] != '\t' {
				start++
			}
			for start < len(line) && (line[start] == ' ' || line[start] == '\t') {
				start++
			}
		}
		args = line[start:]
		result = append(result, process{PID: pid, PPID: ppid, Name: fields[2], Args: args})
	}
	return result
}

func terminate(pid int) error {
	if err := syscall.Kill(pid, syscall.SIGKILL); err != nil {
		return fmt.Errorf("kill: %w", err)
	}
	return nil
}
