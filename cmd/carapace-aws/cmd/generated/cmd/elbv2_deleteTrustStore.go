package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_deleteTrustStoreCmd = &cobra.Command{
	Use:   "delete-trust-store",
	Short: "Deletes a trust store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_deleteTrustStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_deleteTrustStoreCmd).Standalone()

		elbv2_deleteTrustStoreCmd.Flags().String("trust-store-arn", "", "The Amazon Resource Name (ARN) of the trust store.")
		elbv2_deleteTrustStoreCmd.MarkFlagRequired("trust-store-arn")
	})
	elbv2Cmd.AddCommand(elbv2_deleteTrustStoreCmd)
}
