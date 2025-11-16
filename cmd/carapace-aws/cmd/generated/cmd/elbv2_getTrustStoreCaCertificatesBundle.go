package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_getTrustStoreCaCertificatesBundleCmd = &cobra.Command{
	Use:   "get-trust-store-ca-certificates-bundle",
	Short: "Retrieves the ca certificate bundle.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_getTrustStoreCaCertificatesBundleCmd).Standalone()

	elbv2_getTrustStoreCaCertificatesBundleCmd.Flags().String("trust-store-arn", "", "The Amazon Resource Name (ARN) of the trust store.")
	elbv2_getTrustStoreCaCertificatesBundleCmd.MarkFlagRequired("trust-store-arn")
	elbv2Cmd.AddCommand(elbv2_getTrustStoreCaCertificatesBundleCmd)
}
