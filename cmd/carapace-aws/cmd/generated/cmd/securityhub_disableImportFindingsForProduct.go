package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_disableImportFindingsForProductCmd = &cobra.Command{
	Use:   "disable-import-findings-for-product",
	Short: "Disables the integration of the specified product with Security Hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_disableImportFindingsForProductCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_disableImportFindingsForProductCmd).Standalone()

		securityhub_disableImportFindingsForProductCmd.Flags().String("product-subscription-arn", "", "The ARN of the integrated product to disable the integration for.")
		securityhub_disableImportFindingsForProductCmd.MarkFlagRequired("product-subscription-arn")
	})
	securityhubCmd.AddCommand(securityhub_disableImportFindingsForProductCmd)
}
