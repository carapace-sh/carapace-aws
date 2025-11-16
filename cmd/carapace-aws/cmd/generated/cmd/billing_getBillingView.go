package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billing_getBillingViewCmd = &cobra.Command{
	Use:   "get-billing-view",
	Short: "Returns the metadata associated to the specified billing view ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billing_getBillingViewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billing_getBillingViewCmd).Standalone()

		billing_getBillingViewCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) that can be used to uniquely identify the billing view.")
		billing_getBillingViewCmd.MarkFlagRequired("arn")
	})
	billingCmd.AddCommand(billing_getBillingViewCmd)
}
