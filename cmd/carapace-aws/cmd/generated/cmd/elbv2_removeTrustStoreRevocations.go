package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_removeTrustStoreRevocationsCmd = &cobra.Command{
	Use:   "remove-trust-store-revocations",
	Short: "Removes the specified revocation file from the specified trust store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_removeTrustStoreRevocationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_removeTrustStoreRevocationsCmd).Standalone()

		elbv2_removeTrustStoreRevocationsCmd.Flags().String("revocation-ids", "", "The revocation IDs of the revocation files you want to remove.")
		elbv2_removeTrustStoreRevocationsCmd.Flags().String("trust-store-arn", "", "The Amazon Resource Name (ARN) of the trust store.")
		elbv2_removeTrustStoreRevocationsCmd.MarkFlagRequired("revocation-ids")
		elbv2_removeTrustStoreRevocationsCmd.MarkFlagRequired("trust-store-arn")
	})
	elbv2Cmd.AddCommand(elbv2_removeTrustStoreRevocationsCmd)
}
