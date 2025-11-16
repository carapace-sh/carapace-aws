package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_untagCertificateAuthorityCmd = &cobra.Command{
	Use:   "untag-certificate-authority",
	Short: "Remove one or more tags from your private CA.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_untagCertificateAuthorityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(acmPca_untagCertificateAuthorityCmd).Standalone()

		acmPca_untagCertificateAuthorityCmd.Flags().String("certificate-authority-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateCertificateAuthority](https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html).")
		acmPca_untagCertificateAuthorityCmd.Flags().String("tags", "", "List of tags to be removed from the CA.")
		acmPca_untagCertificateAuthorityCmd.MarkFlagRequired("certificate-authority-arn")
		acmPca_untagCertificateAuthorityCmd.MarkFlagRequired("tags")
	})
	acmPcaCmd.AddCommand(acmPca_untagCertificateAuthorityCmd)
}
