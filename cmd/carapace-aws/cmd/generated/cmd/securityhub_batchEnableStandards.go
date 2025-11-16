package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_batchEnableStandardsCmd = &cobra.Command{
	Use:   "batch-enable-standards",
	Short: "Enables the standards specified by the provided `StandardsArn`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_batchEnableStandardsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_batchEnableStandardsCmd).Standalone()

		securityhub_batchEnableStandardsCmd.Flags().String("standards-subscription-requests", "", "The list of standards checks to enable.")
		securityhub_batchEnableStandardsCmd.MarkFlagRequired("standards-subscription-requests")
	})
	securityhubCmd.AddCommand(securityhub_batchEnableStandardsCmd)
}
