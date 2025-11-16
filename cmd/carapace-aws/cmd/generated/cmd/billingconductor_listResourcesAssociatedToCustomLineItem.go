package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_listResourcesAssociatedToCustomLineItemCmd = &cobra.Command{
	Use:   "list-resources-associated-to-custom-line-item",
	Short: "List the resources that are associated to a custom line item.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_listResourcesAssociatedToCustomLineItemCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_listResourcesAssociatedToCustomLineItemCmd).Standalone()

		billingconductor_listResourcesAssociatedToCustomLineItemCmd.Flags().String("arn", "", "The ARN of the custom line item for which the resource associations will be listed.")
		billingconductor_listResourcesAssociatedToCustomLineItemCmd.Flags().String("billing-period", "", "The billing period for which the resource associations will be listed.")
		billingconductor_listResourcesAssociatedToCustomLineItemCmd.Flags().String("filters", "", "(Optional) A `ListResourcesAssociatedToCustomLineItemFilter` that can specify the types of resources that should be retrieved.")
		billingconductor_listResourcesAssociatedToCustomLineItemCmd.Flags().String("max-results", "", "(Optional) The maximum number of resource associations to be retrieved.")
		billingconductor_listResourcesAssociatedToCustomLineItemCmd.Flags().String("next-token", "", "(Optional) The pagination token that's returned by a previous request.")
		billingconductor_listResourcesAssociatedToCustomLineItemCmd.MarkFlagRequired("arn")
	})
	billingconductorCmd.AddCommand(billingconductor_listResourcesAssociatedToCustomLineItemCmd)
}
