package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_addTrustStoreRevocationsCmd = &cobra.Command{
	Use:   "add-trust-store-revocations",
	Short: "Adds the specified revocation file to the specified trust store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_addTrustStoreRevocationsCmd).Standalone()

	elbv2_addTrustStoreRevocationsCmd.Flags().String("revocation-contents", "", "The revocation file to add.")
	elbv2_addTrustStoreRevocationsCmd.Flags().String("trust-store-arn", "", "The Amazon Resource Name (ARN) of the trust store.")
	elbv2_addTrustStoreRevocationsCmd.MarkFlagRequired("trust-store-arn")
	elbv2Cmd.AddCommand(elbv2_addTrustStoreRevocationsCmd)
}
