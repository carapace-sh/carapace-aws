package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billing_deleteBillingViewCmd = &cobra.Command{
	Use:   "delete-billing-view",
	Short: "Deletes the specified billing view.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billing_deleteBillingViewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billing_deleteBillingViewCmd).Standalone()

		billing_deleteBillingViewCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) that can be used to uniquely identify the billing view.")
		billing_deleteBillingViewCmd.Flags().Bool("force", false, "If set to true, forces deletion of the billing view even if it has derived resources (e.g. other billing views or budgets).")
		billing_deleteBillingViewCmd.Flags().Bool("no-force", false, "If set to true, forces deletion of the billing view even if it has derived resources (e.g. other billing views or budgets).")
		billing_deleteBillingViewCmd.MarkFlagRequired("arn")
		billing_deleteBillingViewCmd.Flag("no-force").Hidden = true
	})
	billingCmd.AddCommand(billing_deleteBillingViewCmd)
}
