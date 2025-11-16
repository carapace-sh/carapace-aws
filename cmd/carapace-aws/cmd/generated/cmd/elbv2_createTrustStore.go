package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_createTrustStoreCmd = &cobra.Command{
	Use:   "create-trust-store",
	Short: "Creates a trust store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_createTrustStoreCmd).Standalone()

	elbv2_createTrustStoreCmd.Flags().String("ca-certificates-bundle-s3-bucket", "", "The Amazon S3 bucket for the ca certificates bundle.")
	elbv2_createTrustStoreCmd.Flags().String("ca-certificates-bundle-s3-key", "", "The Amazon S3 path for the ca certificates bundle.")
	elbv2_createTrustStoreCmd.Flags().String("ca-certificates-bundle-s3-object-version", "", "The Amazon S3 object version for the ca certificates bundle.")
	elbv2_createTrustStoreCmd.Flags().String("name", "", "The name of the trust store.")
	elbv2_createTrustStoreCmd.Flags().String("tags", "", "The tags to assign to the trust store.")
	elbv2_createTrustStoreCmd.MarkFlagRequired("ca-certificates-bundle-s3-bucket")
	elbv2_createTrustStoreCmd.MarkFlagRequired("ca-certificates-bundle-s3-key")
	elbv2_createTrustStoreCmd.MarkFlagRequired("name")
	elbv2Cmd.AddCommand(elbv2_createTrustStoreCmd)
}
