package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acm_requestCertificateCmd = &cobra.Command{
	Use:   "request-certificate",
	Short: "Requests an ACM certificate for use with other Amazon Web Services services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acm_requestCertificateCmd).Standalone()

	acm_requestCertificateCmd.Flags().String("certificate-authority-arn", "", "The Amazon Resource Name (ARN) of the private certificate authority (CA) that will be used to issue the certificate.")
	acm_requestCertificateCmd.Flags().String("domain-name", "", "Fully qualified domain name (FQDN), such as www.example.com, that you want to secure with an ACM certificate.")
	acm_requestCertificateCmd.Flags().String("domain-validation-options", "", "The domain name that you want ACM to use to send you emails so that you can validate domain ownership.")
	acm_requestCertificateCmd.Flags().String("idempotency-token", "", "Customer chosen string that can be used to distinguish between calls to `RequestCertificate`.")
	acm_requestCertificateCmd.Flags().String("key-algorithm", "", "Specifies the algorithm of the public and private key pair that your certificate uses to encrypt data.")
	acm_requestCertificateCmd.Flags().String("managed-by", "", "Identifies the Amazon Web Services service that manages the certificate issued by ACM.")
	acm_requestCertificateCmd.Flags().String("options", "", "You can use this parameter to specify whether to add the certificate to a certificate transparency log and export your certificate.")
	acm_requestCertificateCmd.Flags().String("subject-alternative-names", "", "Additional FQDNs to be included in the Subject Alternative Name extension of the ACM certificate.")
	acm_requestCertificateCmd.Flags().String("tags", "", "One or more resource tags to associate with the certificate.")
	acm_requestCertificateCmd.Flags().String("validation-method", "", "The method you want to use if you are requesting a public certificate to validate that you own or control domain.")
	acm_requestCertificateCmd.MarkFlagRequired("domain-name")
	acmCmd.AddCommand(acm_requestCertificateCmd)
}
