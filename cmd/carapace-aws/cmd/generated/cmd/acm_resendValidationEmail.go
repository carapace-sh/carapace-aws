package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acm_resendValidationEmailCmd = &cobra.Command{
	Use:   "resend-validation-email",
	Short: "Resends the email that requests domain ownership validation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acm_resendValidationEmailCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(acm_resendValidationEmailCmd).Standalone()

		acm_resendValidationEmailCmd.Flags().String("certificate-arn", "", "String that contains the ARN of the requested certificate.")
		acm_resendValidationEmailCmd.Flags().String("domain", "", "The fully qualified domain name (FQDN) of the certificate that needs to be validated.")
		acm_resendValidationEmailCmd.Flags().String("validation-domain", "", "The base validation domain that will act as the suffix of the email addresses that are used to send the emails.")
		acm_resendValidationEmailCmd.MarkFlagRequired("certificate-arn")
		acm_resendValidationEmailCmd.MarkFlagRequired("domain")
		acm_resendValidationEmailCmd.MarkFlagRequired("validation-domain")
	})
	acmCmd.AddCommand(acm_resendValidationEmailCmd)
}
