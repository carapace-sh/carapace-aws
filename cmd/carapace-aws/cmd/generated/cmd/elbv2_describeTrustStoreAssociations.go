package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_describeTrustStoreAssociationsCmd = &cobra.Command{
	Use:   "describe-trust-store-associations",
	Short: "Describes all resources associated with the specified trust store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_describeTrustStoreAssociationsCmd).Standalone()

	elbv2_describeTrustStoreAssociationsCmd.Flags().String("marker", "", "The marker for the next set of results.")
	elbv2_describeTrustStoreAssociationsCmd.Flags().String("page-size", "", "The maximum number of results to return with this call.")
	elbv2_describeTrustStoreAssociationsCmd.Flags().String("trust-store-arn", "", "The Amazon Resource Name (ARN) of the trust store.")
	elbv2_describeTrustStoreAssociationsCmd.MarkFlagRequired("trust-store-arn")
	elbv2Cmd.AddCommand(elbv2_describeTrustStoreAssociationsCmd)
}
