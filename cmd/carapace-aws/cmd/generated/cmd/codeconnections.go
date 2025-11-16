package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnectionsCmd = &cobra.Command{
	Use:   "codeconnections",
	Short: "AWS CodeConnections\n\nThis Amazon Web Services CodeConnections API Reference provides descriptions and usage examples of the operations and data types for the Amazon Web Services CodeConnections API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnectionsCmd).Standalone()

	rootCmd.AddCommand(codeconnectionsCmd)
}
