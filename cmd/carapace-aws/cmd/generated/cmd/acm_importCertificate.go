package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acm_importCertificateCmd = &cobra.Command{
	Use:   "import-certificate",
	Short: "Imports a certificate into Certificate Manager (ACM) to use with services that are integrated with ACM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acm_importCertificateCmd).Standalone()

	acm_importCertificateCmd.Flags().String("certificate", "", "The certificate to import.")
	acm_importCertificateCmd.Flags().String("certificate-arn", "", "The [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of an imported certificate to replace.")
	acm_importCertificateCmd.Flags().String("certificate-chain", "", "The PEM encoded certificate chain.")
	acm_importCertificateCmd.Flags().String("private-key", "", "The private key that matches the public key in the certificate.")
	acm_importCertificateCmd.Flags().String("tags", "", "One or more resource tags to associate with the imported certificate.")
	acm_importCertificateCmd.MarkFlagRequired("certificate")
	acm_importCertificateCmd.MarkFlagRequired("private-key")
	acmCmd.AddCommand(acm_importCertificateCmd)
}
