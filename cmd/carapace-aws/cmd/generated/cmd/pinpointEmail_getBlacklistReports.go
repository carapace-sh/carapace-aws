package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_getBlacklistReportsCmd = &cobra.Command{
	Use:   "get-blacklist-reports",
	Short: "Retrieve a list of the blacklists that your dedicated IP addresses appear on.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_getBlacklistReportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointEmail_getBlacklistReportsCmd).Standalone()

		pinpointEmail_getBlacklistReportsCmd.Flags().String("blacklist-item-names", "", "A list of IP addresses that you want to retrieve blacklist information about.")
		pinpointEmail_getBlacklistReportsCmd.MarkFlagRequired("blacklist-item-names")
	})
	pinpointEmailCmd.AddCommand(pinpointEmail_getBlacklistReportsCmd)
}
