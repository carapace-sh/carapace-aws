package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_batchAssociateResourcesToCustomLineItemCmd = &cobra.Command{
	Use:   "batch-associate-resources-to-custom-line-item",
	Short: "Associates a batch of resources to a percentage custom line item.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_batchAssociateResourcesToCustomLineItemCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_batchAssociateResourcesToCustomLineItemCmd).Standalone()

		billingconductor_batchAssociateResourcesToCustomLineItemCmd.Flags().String("billing-period-range", "", "")
		billingconductor_batchAssociateResourcesToCustomLineItemCmd.Flags().String("resource-arns", "", "A list containing the ARNs of the resources to be associated.")
		billingconductor_batchAssociateResourcesToCustomLineItemCmd.Flags().String("target-arn", "", "A percentage custom line item ARN to associate the resources to.")
		billingconductor_batchAssociateResourcesToCustomLineItemCmd.MarkFlagRequired("resource-arns")
		billingconductor_batchAssociateResourcesToCustomLineItemCmd.MarkFlagRequired("target-arn")
	})
	billingconductorCmd.AddCommand(billingconductor_batchAssociateResourcesToCustomLineItemCmd)
}
