package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnectionsCmd = &cobra.Command{
	Use:   "codestar-connections",
	Short: "AWS CodeStar Connections\n\nThis Amazon Web Services CodeStar Connections API Reference provides descriptions and usage examples of the operations and data types for the Amazon Web Services CodeStar Connections API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnectionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codestarConnectionsCmd).Standalone()

	})
	rootCmd.AddCommand(codestarConnectionsCmd)
}
