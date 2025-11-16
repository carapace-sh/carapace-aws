package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_describeProductsCmd = &cobra.Command{
	Use:   "describe-products",
	Short: "Returns information about product integrations in Security Hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_describeProductsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_describeProductsCmd).Standalone()

		securityhub_describeProductsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		securityhub_describeProductsCmd.Flags().String("next-token", "", "The token that is required for pagination.")
		securityhub_describeProductsCmd.Flags().String("product-arn", "", "The ARN of the integration to return.")
	})
	securityhubCmd.AddCommand(securityhub_describeProductsCmd)
}
