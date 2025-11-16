package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_getTrustStoreRevocationContentCmd = &cobra.Command{
	Use:   "get-trust-store-revocation-content",
	Short: "Retrieves the specified revocation file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_getTrustStoreRevocationContentCmd).Standalone()

	elbv2_getTrustStoreRevocationContentCmd.Flags().String("revocation-id", "", "The revocation ID of the revocation file.")
	elbv2_getTrustStoreRevocationContentCmd.Flags().String("trust-store-arn", "", "The Amazon Resource Name (ARN) of the trust store.")
	elbv2_getTrustStoreRevocationContentCmd.MarkFlagRequired("revocation-id")
	elbv2_getTrustStoreRevocationContentCmd.MarkFlagRequired("trust-store-arn")
	elbv2Cmd.AddCommand(elbv2_getTrustStoreRevocationContentCmd)
}
