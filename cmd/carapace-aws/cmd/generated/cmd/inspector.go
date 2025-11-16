package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspectorCmd = &cobra.Command{
	Use:   "inspector",
	Short: "Amazon Inspector\n\nAmazon Inspector enables you to analyze the behavior of your AWS resources and to identify potential security issues.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspectorCmd).Standalone()

	rootCmd.AddCommand(inspectorCmd)
}
