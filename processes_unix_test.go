//go:build !windows

package main

import (
	"reflect"
	"testing"
)

func TestParsePSOutput(t *testing.T) {
	output := "  101   1 node /usr/bin/node  /opt/@openai/codex/bin/codex\n" +
		"  202 101 python3 python3 -m aider_chat\n" +
		"invalid row\n"
	want := []process{
		{PID: 101, PPID: 1, Name: "node", Args: "/usr/bin/node  /opt/@openai/codex/bin/codex"},
		{PID: 202, PPID: 101, Name: "python3", Args: "python3 -m aider_chat"},
	}
	if got := parsePSOutput(output); !reflect.DeepEqual(got, want) {
		t.Fatalf("parsePSOutput() = %#v, want %#v", got, want)
	}
}
