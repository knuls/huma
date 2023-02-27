package cmd_test

import (
	"testing"

	"github.com/knuls/huma/cmd"
)

func TestOdinCommand(t *testing.T) {
	cmd := cmd.NewOdinCommand()
	cmd.Execute()
}
