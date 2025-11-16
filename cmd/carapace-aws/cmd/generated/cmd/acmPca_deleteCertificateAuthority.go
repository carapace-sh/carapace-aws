package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_deleteCertificateAuthorityCmd = &cobra.Command{
	Use:   "delete-certificate-authority",
	Short: "Deletes a private certificate authority (CA).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_deleteCertificateAuthorityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(acmPca_deleteCertificateAuthorityCmd).Standalone()

		acmPca_deleteCertificateAuthorityCmd.Flags().String("certificate-authority-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateCertificateAuthority](https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html).")
		acmPca_deleteCertificateAuthorityCmd.Flags().String("permanent-deletion-time-in-days", "", "The number of days to make a CA restorable after it has been deleted.")
		acmPca_deleteCertificateAuthorityCmd.MarkFlagRequired("certificate-authority-arn")
	})
	acmPcaCmd.AddCommand(acmPca_deleteCertificateAuthorityCmd)
}
