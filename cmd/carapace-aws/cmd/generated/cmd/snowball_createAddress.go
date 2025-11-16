package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_createAddressCmd = &cobra.Command{
	Use:   "create-address",
	Short: "Creates an address for a Snow device to be shipped to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_createAddressCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snowball_createAddressCmd).Standalone()

		snowball_createAddressCmd.Flags().String("address", "", "The address that you want the Snow device shipped to.")
		snowball_createAddressCmd.MarkFlagRequired("address")
	})
	snowballCmd.AddCommand(snowball_createAddressCmd)
}
