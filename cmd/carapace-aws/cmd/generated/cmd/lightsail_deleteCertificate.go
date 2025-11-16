package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteCertificateCmd = &cobra.Command{
	Use:   "delete-certificate",
	Short: "Deletes an SSL/TLS certificate for your Amazon Lightsail content delivery network (CDN) distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_deleteCertificateCmd).Standalone()

		lightsail_deleteCertificateCmd.Flags().String("certificate-name", "", "The name of the certificate to delete.")
		lightsail_deleteCertificateCmd.MarkFlagRequired("certificate-name")
	})
	lightsailCmd.AddCommand(lightsail_deleteCertificateCmd)
}
