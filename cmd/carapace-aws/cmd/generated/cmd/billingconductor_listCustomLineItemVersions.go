package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_listCustomLineItemVersionsCmd = &cobra.Command{
	Use:   "list-custom-line-item-versions",
	Short: "A paginated call to get a list of all custom line item versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_listCustomLineItemVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_listCustomLineItemVersionsCmd).Standalone()

		billingconductor_listCustomLineItemVersionsCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) for the custom line item.")
		billingconductor_listCustomLineItemVersionsCmd.Flags().String("filters", "", "A `ListCustomLineItemVersionsFilter` that specifies the billing period range in which the custom line item versions are applied.")
		billingconductor_listCustomLineItemVersionsCmd.Flags().String("max-results", "", "The maximum number of custom line item versions to retrieve.")
		billingconductor_listCustomLineItemVersionsCmd.Flags().String("next-token", "", "The pagination token that's used on subsequent calls to retrieve custom line item versions.")
		billingconductor_listCustomLineItemVersionsCmd.MarkFlagRequired("arn")
	})
	billingconductorCmd.AddCommand(billingconductor_listCustomLineItemVersionsCmd)
}
