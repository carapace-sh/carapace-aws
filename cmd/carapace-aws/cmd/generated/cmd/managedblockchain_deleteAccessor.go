package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_deleteAccessorCmd = &cobra.Command{
	Use:   "delete-accessor",
	Short: "Deletes an accessor that your Amazon Web Services account owns.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_deleteAccessorCmd).Standalone()

	managedblockchain_deleteAccessorCmd.Flags().String("accessor-id", "", "The unique identifier of the accessor.")
	managedblockchain_deleteAccessorCmd.MarkFlagRequired("accessor-id")
	managedblockchainCmd.AddCommand(managedblockchain_deleteAccessorCmd)
}
