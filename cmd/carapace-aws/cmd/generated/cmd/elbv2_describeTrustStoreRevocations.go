package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_describeTrustStoreRevocationsCmd = &cobra.Command{
	Use:   "describe-trust-store-revocations",
	Short: "Describes the revocation files in use by the specified trust store or revocation files.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_describeTrustStoreRevocationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_describeTrustStoreRevocationsCmd).Standalone()

		elbv2_describeTrustStoreRevocationsCmd.Flags().String("marker", "", "The marker for the next set of results.")
		elbv2_describeTrustStoreRevocationsCmd.Flags().String("page-size", "", "The maximum number of results to return with this call.")
		elbv2_describeTrustStoreRevocationsCmd.Flags().String("revocation-ids", "", "The revocation IDs of the revocation files you want to describe.")
		elbv2_describeTrustStoreRevocationsCmd.Flags().String("trust-store-arn", "", "The Amazon Resource Name (ARN) of the trust store.")
		elbv2_describeTrustStoreRevocationsCmd.MarkFlagRequired("trust-store-arn")
	})
	elbv2Cmd.AddCommand(elbv2_describeTrustStoreRevocationsCmd)
}
