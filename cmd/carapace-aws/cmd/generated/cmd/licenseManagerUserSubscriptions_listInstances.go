package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var licenseManagerUserSubscriptions_listInstancesCmd = &cobra.Command{
	Use:   "list-instances",
	Short: "Lists the EC2 instances providing user-based subscriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(licenseManagerUserSubscriptions_listInstancesCmd).Standalone()

	licenseManagerUserSubscriptions_listInstancesCmd.Flags().String("filters", "", "You can use the following filters to streamline results:")
	licenseManagerUserSubscriptions_listInstancesCmd.Flags().String("max-results", "", "The maximum number of results to return from a single request.")
	licenseManagerUserSubscriptions_listInstancesCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	licenseManagerUserSubscriptionsCmd.AddCommand(licenseManagerUserSubscriptions_listInstancesCmd)
}
