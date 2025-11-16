package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_uploadSigningCertificateCmd = &cobra.Command{
	Use:   "upload-signing-certificate",
	Short: "Uploads an X.509 signing certificate and associates it with the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_uploadSigningCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_uploadSigningCertificateCmd).Standalone()

		iam_uploadSigningCertificateCmd.Flags().String("certificate-body", "", "The contents of the signing certificate.")
		iam_uploadSigningCertificateCmd.Flags().String("user-name", "", "The name of the user the signing certificate is for.")
		iam_uploadSigningCertificateCmd.MarkFlagRequired("certificate-body")
	})
	iamCmd.AddCommand(iam_uploadSigningCertificateCmd)
}
