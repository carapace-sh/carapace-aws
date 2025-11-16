package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_updateCertificateCmd = &cobra.Command{
	Use:   "update-certificate",
	Short: "Updates the active and inactive dates for a certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_updateCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_updateCertificateCmd).Standalone()

		transfer_updateCertificateCmd.Flags().String("active-date", "", "An optional date that specifies when the certificate becomes active.")
		transfer_updateCertificateCmd.Flags().String("certificate-id", "", "The identifier of the certificate object that you are updating.")
		transfer_updateCertificateCmd.Flags().String("description", "", "A short description to help identify the certificate.")
		transfer_updateCertificateCmd.Flags().String("inactive-date", "", "An optional date that specifies when the certificate becomes inactive.")
		transfer_updateCertificateCmd.MarkFlagRequired("certificate-id")
	})
	transferCmd.AddCommand(transfer_updateCertificateCmd)
}
