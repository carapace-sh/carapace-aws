package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_describeTrustStoresCmd = &cobra.Command{
	Use:   "describe-trust-stores",
	Short: "Describes all trust stores for the specified account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_describeTrustStoresCmd).Standalone()

	elbv2_describeTrustStoresCmd.Flags().String("marker", "", "The marker for the next set of results.")
	elbv2_describeTrustStoresCmd.Flags().String("names", "", "The names of the trust stores.")
	elbv2_describeTrustStoresCmd.Flags().String("page-size", "", "The maximum number of results to return with this call.")
	elbv2_describeTrustStoresCmd.Flags().String("trust-store-arns", "", "The Amazon Resource Name (ARN) of the trust store.")
	elbv2Cmd.AddCommand(elbv2_describeTrustStoresCmd)
}
