package cmd

import (
	"github.com/knuls/huma/internal/hugin"
	"github.com/spf13/cobra"
)

func NewHuginCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "hugin",
		Short: "The hugin service serves and processes actions.",
		Run: func(cmd *cobra.Command, args []string) {
			hugin.New()
		},
	}
}

func init() {
	rootCmd.AddCommand(NewHuginCommand())
}
