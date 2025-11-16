package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerLinuxSubscriptions_listLinuxSubscriptionsCmd = &cobra.Command{
	Use:   "list-linux-subscriptions",
	Short: "Lists the Linux subscriptions that have been discovered.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerLinuxSubscriptions_listLinuxSubscriptionsCmd).Standalone()

	licenseManagerLinuxSubscriptions_listLinuxSubscriptionsCmd.Flags().String("filters", "", "An array of structures that you can use to filter the results to those that match one or more sets of key-value pairs that you specify.")
	licenseManagerLinuxSubscriptions_listLinuxSubscriptionsCmd.Flags().String("max-results", "", "The maximum items to return in a request.")
	licenseManagerLinuxSubscriptions_listLinuxSubscriptionsCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	licenseManagerLinuxSubscriptionsCmd.AddCommand(licenseManagerLinuxSubscriptions_listLinuxSubscriptionsCmd)
}
