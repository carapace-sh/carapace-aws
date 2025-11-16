package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_listEnabledProductsForImportCmd = &cobra.Command{
	Use:   "list-enabled-products-for-import",
	Short: "Lists all findings-generating solutions (products) that you are subscribed to receive findings from in Security Hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_listEnabledProductsForImportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_listEnabledProductsForImportCmd).Standalone()

		securityhub_listEnabledProductsForImportCmd.Flags().String("max-results", "", "The maximum number of items to return in the response.")
		securityhub_listEnabledProductsForImportCmd.Flags().String("next-token", "", "The token that is required for pagination.")
	})
	securityhubCmd.AddCommand(securityhub_listEnabledProductsForImportCmd)
}
