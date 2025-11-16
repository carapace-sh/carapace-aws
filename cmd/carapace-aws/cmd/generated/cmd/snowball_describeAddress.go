package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_describeAddressCmd = &cobra.Command{
	Use:   "describe-address",
	Short: "Takes an `AddressId` and returns specific details about that address in the form of an `Address` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_describeAddressCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(snowball_describeAddressCmd).Standalone()

		snowball_describeAddressCmd.Flags().String("address-id", "", "The automatically generated ID for a specific address.")
		snowball_describeAddressCmd.MarkFlagRequired("address-id")
	})
	snowballCmd.AddCommand(snowball_describeAddressCmd)
}
