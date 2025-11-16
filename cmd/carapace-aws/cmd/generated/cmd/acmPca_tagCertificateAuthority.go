package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_tagCertificateAuthorityCmd = &cobra.Command{
	Use:   "tag-certificate-authority",
	Short: "Adds one or more tags to your private CA.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_tagCertificateAuthorityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(acmPca_tagCertificateAuthorityCmd).Standalone()

		acmPca_tagCertificateAuthorityCmd.Flags().String("certificate-authority-arn", "", "The Amazon Resource Name (ARN) that was returned when you called [CreateCertificateAuthority](https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html).")
		acmPca_tagCertificateAuthorityCmd.Flags().String("tags", "", "List of tags to be associated with the CA.")
		acmPca_tagCertificateAuthorityCmd.MarkFlagRequired("certificate-authority-arn")
		acmPca_tagCertificateAuthorityCmd.MarkFlagRequired("tags")
	})
	acmPcaCmd.AddCommand(acmPca_tagCertificateAuthorityCmd)
}
