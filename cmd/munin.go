package cmd

import (
	"github.com/knuls/huma/internal/munin"
	"github.com/spf13/cobra"
)

func NewMuninCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "munin",
		Short: "The munin service serves and processes pipelines.",
		Run: func(cmd *cobra.Command, args []string) {
			munin.New()
		},
	}
}

func init() {
	rootCmd.AddCommand(NewMuninCommand())
}
