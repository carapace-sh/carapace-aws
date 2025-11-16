package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerLinuxSubscriptions_listLinuxSubscriptionInstancesCmd = &cobra.Command{
	Use:   "list-linux-subscription-instances",
	Short: "Lists the running Amazon EC2 instances that were discovered with commercial Linux subscriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerLinuxSubscriptions_listLinuxSubscriptionInstancesCmd).Standalone()

	licenseManagerLinuxSubscriptions_listLinuxSubscriptionInstancesCmd.Flags().String("filters", "", "An array of structures that you can use to filter the results by your specified criteria.")
	licenseManagerLinuxSubscriptions_listLinuxSubscriptionInstancesCmd.Flags().String("max-results", "", "The maximum items to return in a request.")
	licenseManagerLinuxSubscriptions_listLinuxSubscriptionInstancesCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	licenseManagerLinuxSubscriptionsCmd.AddCommand(licenseManagerLinuxSubscriptions_listLinuxSubscriptionInstancesCmd)
}
