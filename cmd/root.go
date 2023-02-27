package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "huma",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
