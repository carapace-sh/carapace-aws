package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_attachCertificateToDistributionCmd = &cobra.Command{
	Use:   "attach-certificate-to-distribution",
	Short: "Attaches an SSL/TLS certificate to your Amazon Lightsail content delivery network (CDN) distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_attachCertificateToDistributionCmd).Standalone()

	lightsail_attachCertificateToDistributionCmd.Flags().String("certificate-name", "", "The name of the certificate to attach to a distribution.")
	lightsail_attachCertificateToDistributionCmd.Flags().String("distribution-name", "", "The name of the distribution that the certificate will be attached to.")
	lightsail_attachCertificateToDistributionCmd.MarkFlagRequired("certificate-name")
	lightsail_attachCertificateToDistributionCmd.MarkFlagRequired("distribution-name")
	lightsailCmd.AddCommand(lightsail_attachCertificateToDistributionCmd)
}
