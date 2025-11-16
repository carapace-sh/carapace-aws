package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchainQueryCmd = &cobra.Command{
	Use:   "managedblockchain-query",
	Short: "Amazon Managed Blockchain (AMB) Query provides you with convenient access to multi-blockchain network data, which makes it easier for you to extract contextual data related to blockchain activity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchainQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(managedblockchainQueryCmd).Standalone()

	})
	rootCmd.AddCommand(managedblockchainQueryCmd)
}
