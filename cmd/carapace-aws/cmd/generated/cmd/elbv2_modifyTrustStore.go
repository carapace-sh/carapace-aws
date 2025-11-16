package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_modifyTrustStoreCmd = &cobra.Command{
	Use:   "modify-trust-store",
	Short: "Update the ca certificate bundle for the specified trust store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_modifyTrustStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_modifyTrustStoreCmd).Standalone()

		elbv2_modifyTrustStoreCmd.Flags().String("ca-certificates-bundle-s3-bucket", "", "The Amazon S3 bucket for the ca certificates bundle.")
		elbv2_modifyTrustStoreCmd.Flags().String("ca-certificates-bundle-s3-key", "", "The Amazon S3 path for the ca certificates bundle.")
		elbv2_modifyTrustStoreCmd.Flags().String("ca-certificates-bundle-s3-object-version", "", "The Amazon S3 object version for the ca certificates bundle.")
		elbv2_modifyTrustStoreCmd.Flags().String("trust-store-arn", "", "The Amazon Resource Name (ARN) of the trust store.")
		elbv2_modifyTrustStoreCmd.MarkFlagRequired("ca-certificates-bundle-s3-bucket")
		elbv2_modifyTrustStoreCmd.MarkFlagRequired("ca-certificates-bundle-s3-key")
		elbv2_modifyTrustStoreCmd.MarkFlagRequired("trust-store-arn")
	})
	elbv2Cmd.AddCommand(elbv2_modifyTrustStoreCmd)
}
