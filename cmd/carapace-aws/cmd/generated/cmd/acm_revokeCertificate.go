package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acm_revokeCertificateCmd = &cobra.Command{
	Use:   "revoke-certificate",
	Short: "Revokes a public ACM certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acm_revokeCertificateCmd).Standalone()

	acm_revokeCertificateCmd.Flags().String("certificate-arn", "", "The Amazon Resource Name (ARN) of the public or private certificate that will be revoked.")
	acm_revokeCertificateCmd.Flags().String("revocation-reason", "", "Specifies why you revoked the certificate.")
	acm_revokeCertificateCmd.MarkFlagRequired("certificate-arn")
	acm_revokeCertificateCmd.MarkFlagRequired("revocation-reason")
	acmCmd.AddCommand(acm_revokeCertificateCmd)
}
