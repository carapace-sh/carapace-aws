package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_listAccessorsCmd = &cobra.Command{
	Use:   "list-accessors",
	Short: "Returns a list of the accessors and their properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_listAccessorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(managedblockchain_listAccessorsCmd).Standalone()

		managedblockchain_listAccessorsCmd.Flags().String("max-results", "", "The maximum number of accessors to list.")
		managedblockchain_listAccessorsCmd.Flags().String("network-type", "", "The blockchain network that the `Accessor` token is created for.")
		managedblockchain_listAccessorsCmd.Flags().String("next-token", "", "The pagination token that indicates the next set of results to retrieve.")
	})
	managedblockchainCmd.AddCommand(managedblockchain_listAccessorsCmd)
}
