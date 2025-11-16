package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getKeyPairsCmd = &cobra.Command{
	Use:   "get-key-pairs",
	Short: "Returns information about all key pairs in the user's account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getKeyPairsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getKeyPairsCmd).Standalone()

		lightsail_getKeyPairsCmd.Flags().String("include-default-key-pair", "", "A Boolean value that indicates whether to include the default key pair in the response of your request.")
		lightsail_getKeyPairsCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	})
	lightsailCmd.AddCommand(lightsail_getKeyPairsCmd)
}
