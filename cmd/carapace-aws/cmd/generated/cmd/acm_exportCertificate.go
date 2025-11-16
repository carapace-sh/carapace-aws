package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acm_exportCertificateCmd = &cobra.Command{
	Use:   "export-certificate",
	Short: "Exports a private certificate issued by a private certificate authority (CA) or public certificate for use anywhere.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acm_exportCertificateCmd).Standalone()

	acm_exportCertificateCmd.Flags().String("certificate-arn", "", "An Amazon Resource Name (ARN) of the issued certificate.")
	acm_exportCertificateCmd.Flags().String("passphrase", "", "Passphrase to associate with the encrypted exported private key.")
	acm_exportCertificateCmd.MarkFlagRequired("certificate-arn")
	acm_exportCertificateCmd.MarkFlagRequired("passphrase")
	acmCmd.AddCommand(acm_exportCertificateCmd)
}
