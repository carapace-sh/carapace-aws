package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_detachCertificateFromDistributionCmd = &cobra.Command{
	Use:   "detach-certificate-from-distribution",
	Short: "Detaches an SSL/TLS certificate from your Amazon Lightsail content delivery network (CDN) distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_detachCertificateFromDistributionCmd).Standalone()

	lightsail_detachCertificateFromDistributionCmd.Flags().String("distribution-name", "", "The name of the distribution from which to detach the certificate.")
	lightsail_detachCertificateFromDistributionCmd.MarkFlagRequired("distribution-name")
	lightsailCmd.AddCommand(lightsail_detachCertificateFromDistributionCmd)
}
