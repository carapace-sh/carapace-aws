package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_describeCertificateCmd = &cobra.Command{
	Use:   "describe-certificate",
	Short: "Describes the certificate that's identified by the `CertificateId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_describeCertificateCmd).Standalone()

	transfer_describeCertificateCmd.Flags().String("certificate-id", "", "An array of identifiers for the imported certificates.")
	transfer_describeCertificateCmd.MarkFlagRequired("certificate-id")
	transferCmd.AddCommand(transfer_describeCertificateCmd)
}
