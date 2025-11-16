package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_importCertificateCmd = &cobra.Command{
	Use:   "import-certificate",
	Short: "Uploads the specified certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_importCertificateCmd).Standalone()

	dms_importCertificateCmd.Flags().String("certificate-identifier", "", "A customer-assigned name for the certificate.")
	dms_importCertificateCmd.Flags().String("certificate-pem", "", "The contents of a `.pem` file, which contains an X.509 certificate.")
	dms_importCertificateCmd.Flags().String("certificate-wallet", "", "The location of an imported Oracle Wallet certificate for use with SSL.")
	dms_importCertificateCmd.Flags().String("tags", "", "The tags associated with the certificate.")
	dms_importCertificateCmd.MarkFlagRequired("certificate-identifier")
	dmsCmd.AddCommand(dms_importCertificateCmd)
}
