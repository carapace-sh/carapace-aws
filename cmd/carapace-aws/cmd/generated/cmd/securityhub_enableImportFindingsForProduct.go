package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_enableImportFindingsForProductCmd = &cobra.Command{
	Use:   "enable-import-findings-for-product",
	Short: "Enables the integration of a partner product with Security Hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_enableImportFindingsForProductCmd).Standalone()

	securityhub_enableImportFindingsForProductCmd.Flags().String("product-arn", "", "The ARN of the product to enable the integration for.")
	securityhub_enableImportFindingsForProductCmd.MarkFlagRequired("product-arn")
	securityhubCmd.AddCommand(securityhub_enableImportFindingsForProductCmd)
}
