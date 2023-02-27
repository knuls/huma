package cmd

import (
	"github.com/knuls/huma/internal/odin"
	"github.com/spf13/cobra"
)

func NewOdinCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "odin",
		Short: "The odin service.",
		Long:  "The odin service serves the platform across creators and organizations.",
		Run: func(cmd *cobra.Command, args []string) {
			odin.New()
		},
	}
}

func init() {
	rootCmd.AddCommand(NewOdinCommand())
}
