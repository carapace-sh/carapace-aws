package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_listBillingGroupsCmd = &cobra.Command{
	Use:   "list-billing-groups",
	Short: "A paginated call to retrieve a list of billing groups for the given billing period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_listBillingGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_listBillingGroupsCmd).Standalone()

		billingconductor_listBillingGroupsCmd.Flags().String("billing-period", "", "The preferred billing period to get billing groups.")
		billingconductor_listBillingGroupsCmd.Flags().String("filters", "", "A `ListBillingGroupsFilter` that specifies the billing group and pricing plan to retrieve billing group information.")
		billingconductor_listBillingGroupsCmd.Flags().String("max-results", "", "The maximum number of billing groups to retrieve.")
		billingconductor_listBillingGroupsCmd.Flags().String("next-token", "", "The pagination token that's used on subsequent calls to get billing groups.")
	})
	billingconductorCmd.AddCommand(billingconductor_listBillingGroupsCmd)
}
