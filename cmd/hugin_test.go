package cmd_test

import (
	"testing"

	"github.com/knuls/huma/cmd"
)

func TestHuginCommand(t *testing.T) {
	cmd := cmd.NewHuginCommand()
	cmd.Execute()
}
