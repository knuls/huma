package cmd_test

import (
	"testing"

	"github.com/knuls/huma/cmd"
)

func TestMuninCommand(t *testing.T) {
	cmd := cmd.NewMuninCommand()
	cmd.Execute()
}
