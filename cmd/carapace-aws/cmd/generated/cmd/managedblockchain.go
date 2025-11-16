package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchainCmd = &cobra.Command{
	Use:   "managedblockchain",
	Short: "Amazon Managed Blockchain is a fully managed service for creating and managing blockchain networks using open-source frameworks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(managedblockchainCmd).Standalone()

	})
	rootCmd.AddCommand(managedblockchainCmd)
}
