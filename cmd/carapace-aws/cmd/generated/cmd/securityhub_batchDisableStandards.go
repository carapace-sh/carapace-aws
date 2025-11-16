package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_batchDisableStandardsCmd = &cobra.Command{
	Use:   "batch-disable-standards",
	Short: "Disables the standards specified by the provided `StandardsSubscriptionArns`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_batchDisableStandardsCmd).Standalone()

	securityhub_batchDisableStandardsCmd.Flags().String("standards-subscription-arns", "", "The ARNs of the standards subscriptions to disable.")
	securityhub_batchDisableStandardsCmd.MarkFlagRequired("standards-subscription-arns")
	securityhubCmd.AddCommand(securityhub_batchDisableStandardsCmd)
}
