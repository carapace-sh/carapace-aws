package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_listBillingGroupCostReportsCmd = &cobra.Command{
	Use:   "list-billing-group-cost-reports",
	Short: "A paginated call to retrieve a summary report of actual Amazon Web Services charges and the calculated Amazon Web Services charges based on the associated pricing plan of a billing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_listBillingGroupCostReportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_listBillingGroupCostReportsCmd).Standalone()

		billingconductor_listBillingGroupCostReportsCmd.Flags().String("billing-period", "", "The preferred billing period for your report.")
		billingconductor_listBillingGroupCostReportsCmd.Flags().String("filters", "", "A `ListBillingGroupCostReportsFilter` to specify billing groups to retrieve reports from.")
		billingconductor_listBillingGroupCostReportsCmd.Flags().String("max-results", "", "The maximum number of reports to retrieve.")
		billingconductor_listBillingGroupCostReportsCmd.Flags().String("next-token", "", "The pagination token that's used on subsequent calls to get reports.")
	})
	billingconductorCmd.AddCommand(billingconductor_listBillingGroupCostReportsCmd)
}
