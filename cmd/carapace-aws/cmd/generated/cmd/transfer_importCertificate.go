package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_importCertificateCmd = &cobra.Command{
	Use:   "import-certificate",
	Short: "Imports the signing and encryption certificates that you need to create local (AS2) profiles and partner profiles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_importCertificateCmd).Standalone()

	transfer_importCertificateCmd.Flags().String("active-date", "", "An optional date that specifies when the certificate becomes active.")
	transfer_importCertificateCmd.Flags().String("certificate", "", "- For the CLI, provide a file path for a certificate in URI format.")
	transfer_importCertificateCmd.Flags().String("certificate-chain", "", "An optional list of certificates that make up the chain for the certificate that's being imported.")
	transfer_importCertificateCmd.Flags().String("description", "", "A short description that helps identify the certificate.")
	transfer_importCertificateCmd.Flags().String("inactive-date", "", "An optional date that specifies when the certificate becomes inactive.")
	transfer_importCertificateCmd.Flags().String("private-key", "", "- For the CLI, provide a file path for a private key in URI format.")
	transfer_importCertificateCmd.Flags().String("tags", "", "Key-value pairs that can be used to group and search for certificates.")
	transfer_importCertificateCmd.Flags().String("usage", "", "Specifies how this certificate is used.")
	transfer_importCertificateCmd.MarkFlagRequired("certificate")
	transfer_importCertificateCmd.MarkFlagRequired("usage")
	transferCmd.AddCommand(transfer_importCertificateCmd)
}
