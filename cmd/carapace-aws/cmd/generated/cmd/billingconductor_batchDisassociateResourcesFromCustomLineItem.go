package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_batchDisassociateResourcesFromCustomLineItemCmd = &cobra.Command{
	Use:   "batch-disassociate-resources-from-custom-line-item",
	Short: "Disassociates a batch of resources from a percentage custom line item.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_batchDisassociateResourcesFromCustomLineItemCmd).Standalone()

	billingconductor_batchDisassociateResourcesFromCustomLineItemCmd.Flags().String("billing-period-range", "", "")
	billingconductor_batchDisassociateResourcesFromCustomLineItemCmd.Flags().String("resource-arns", "", "A list containing the ARNs of resources to be disassociated.")
	billingconductor_batchDisassociateResourcesFromCustomLineItemCmd.Flags().String("target-arn", "", "A percentage custom line item ARN to disassociate the resources from.")
	billingconductor_batchDisassociateResourcesFromCustomLineItemCmd.MarkFlagRequired("resource-arns")
	billingconductor_batchDisassociateResourcesFromCustomLineItemCmd.MarkFlagRequired("target-arn")
	billingconductorCmd.AddCommand(billingconductor_batchDisassociateResourcesFromCustomLineItemCmd)
}
