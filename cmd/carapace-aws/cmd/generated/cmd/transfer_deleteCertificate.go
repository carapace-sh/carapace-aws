package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_deleteCertificateCmd = &cobra.Command{
	Use:   "delete-certificate",
	Short: "Deletes the certificate that's specified in the `CertificateId` parameter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_deleteCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_deleteCertificateCmd).Standalone()

		transfer_deleteCertificateCmd.Flags().String("certificate-id", "", "The identifier of the certificate object that you are deleting.")
		transfer_deleteCertificateCmd.MarkFlagRequired("certificate-id")
	})
	transferCmd.AddCommand(transfer_deleteCertificateCmd)
}
