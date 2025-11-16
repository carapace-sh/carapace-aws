package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_getAccessorCmd = &cobra.Command{
	Use:   "get-accessor",
	Short: "Returns detailed information about an accessor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_getAccessorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(managedblockchain_getAccessorCmd).Standalone()

		managedblockchain_getAccessorCmd.Flags().String("accessor-id", "", "The unique identifier of the accessor.")
		managedblockchain_getAccessorCmd.MarkFlagRequired("accessor-id")
	})
	managedblockchainCmd.AddCommand(managedblockchain_getAccessorCmd)
}
