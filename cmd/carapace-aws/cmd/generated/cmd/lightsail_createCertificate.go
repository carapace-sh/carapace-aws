package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createCertificateCmd = &cobra.Command{
	Use:   "create-certificate",
	Short: "Creates an SSL/TLS certificate for an Amazon Lightsail content delivery network (CDN) distribution and a container service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_createCertificateCmd).Standalone()

		lightsail_createCertificateCmd.Flags().String("certificate-name", "", "The name for the certificate.")
		lightsail_createCertificateCmd.Flags().String("domain-name", "", "The domain name (`example.com`) for the certificate.")
		lightsail_createCertificateCmd.Flags().String("subject-alternative-names", "", "An array of strings that specify the alternate domains (`example2.com`) and subdomains (`blog.example.com`) for the certificate.")
		lightsail_createCertificateCmd.Flags().String("tags", "", "The tag keys and optional values to add to the certificate during create.")
		lightsail_createCertificateCmd.MarkFlagRequired("certificate-name")
		lightsail_createCertificateCmd.MarkFlagRequired("domain-name")
	})
	lightsailCmd.AddCommand(lightsail_createCertificateCmd)
}
