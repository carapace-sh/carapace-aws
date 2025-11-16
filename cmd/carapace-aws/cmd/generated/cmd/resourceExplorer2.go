package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2Cmd = &cobra.Command{
	Use:   "resource-explorer-2",
	Short: "Amazon Web Services Resource Explorer is a resource search and discovery service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2Cmd).Standalone()

	rootCmd.AddCommand(resourceExplorer2Cmd)
}
