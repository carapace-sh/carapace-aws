package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_uploadServerCertificateCmd = &cobra.Command{
	Use:   "upload-server-certificate",
	Short: "Uploads a server certificate entity for the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_uploadServerCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_uploadServerCertificateCmd).Standalone()

		iam_uploadServerCertificateCmd.Flags().String("certificate-body", "", "The contents of the public key certificate in PEM-encoded format.")
		iam_uploadServerCertificateCmd.Flags().String("certificate-chain", "", "The contents of the certificate chain.")
		iam_uploadServerCertificateCmd.Flags().String("path", "", "The path for the server certificate.")
		iam_uploadServerCertificateCmd.Flags().String("private-key", "", "The contents of the private key in PEM-encoded format.")
		iam_uploadServerCertificateCmd.Flags().String("server-certificate-name", "", "The name for the server certificate.")
		iam_uploadServerCertificateCmd.Flags().String("tags", "", "A list of tags that you want to attach to the new IAM server certificate resource.")
		iam_uploadServerCertificateCmd.MarkFlagRequired("certificate-body")
		iam_uploadServerCertificateCmd.MarkFlagRequired("private-key")
		iam_uploadServerCertificateCmd.MarkFlagRequired("server-certificate-name")
	})
	iamCmd.AddCommand(iam_uploadServerCertificateCmd)
}
