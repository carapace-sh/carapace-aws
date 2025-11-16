package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getCertificatesCmd = &cobra.Command{
	Use:   "get-certificates",
	Short: "Returns information about one or more Amazon Lightsail SSL/TLS certificates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getCertificatesCmd).Standalone()

	lightsail_getCertificatesCmd.Flags().String("certificate-name", "", "The name for the certificate for which to return information.")
	lightsail_getCertificatesCmd.Flags().String("certificate-statuses", "", "The status of the certificates for which to return information.")
	lightsail_getCertificatesCmd.Flags().String("include-certificate-details", "", "Indicates whether to include detailed information about the certificates in the response.")
	lightsail_getCertificatesCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	lightsailCmd.AddCommand(lightsail_getCertificatesCmd)
}
