package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_describeAddressesCmd = &cobra.Command{
	Use:   "describe-addresses",
	Short: "Returns a specified number of `ADDRESS` objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_describeAddressesCmd).Standalone()

	snowball_describeAddressesCmd.Flags().String("max-results", "", "The number of `ADDRESS` objects to return.")
	snowball_describeAddressesCmd.Flags().String("next-token", "", "HTTP requests are stateless.")
	snowballCmd.AddCommand(snowball_describeAddressesCmd)
}
