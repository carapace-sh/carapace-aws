package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_getBillingGroupCostReportCmd = &cobra.Command{
	Use:   "get-billing-group-cost-report",
	Short: "Retrieves the margin summary report, which includes the Amazon Web Services cost and charged amount (pro forma cost) by Amazon Web Services service for a specific billing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_getBillingGroupCostReportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_getBillingGroupCostReportCmd).Standalone()

		billingconductor_getBillingGroupCostReportCmd.Flags().String("arn", "", "The Amazon Resource Number (ARN) that uniquely identifies the billing group.")
		billingconductor_getBillingGroupCostReportCmd.Flags().String("billing-period-range", "", "A time range for which the margin summary is effective.")
		billingconductor_getBillingGroupCostReportCmd.Flags().String("group-by", "", "A list of strings that specify the attributes that are used to break down costs in the margin summary reports for the billing group.")
		billingconductor_getBillingGroupCostReportCmd.Flags().String("max-results", "", "The maximum number of margin summary reports to retrieve.")
		billingconductor_getBillingGroupCostReportCmd.Flags().String("next-token", "", "The pagination token used on subsequent calls to get reports.")
		billingconductor_getBillingGroupCostReportCmd.MarkFlagRequired("arn")
	})
	billingconductorCmd.AddCommand(billingconductor_getBillingGroupCostReportCmd)
}
