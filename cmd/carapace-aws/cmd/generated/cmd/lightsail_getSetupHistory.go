package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getSetupHistoryCmd = &cobra.Command{
	Use:   "get-setup-history",
	Short: "Returns detailed information for five of the most recent `SetupInstanceHttps` requests that were ran on the target instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getSetupHistoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getSetupHistoryCmd).Standalone()

		lightsail_getSetupHistoryCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
		lightsail_getSetupHistoryCmd.Flags().String("resource-name", "", "The name of the resource for which you are requesting information.")
		lightsail_getSetupHistoryCmd.MarkFlagRequired("resource-name")
	})
	lightsailCmd.AddCommand(lightsail_getSetupHistoryCmd)
}
