package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acm_renewCertificateCmd = &cobra.Command{
	Use:   "renew-certificate",
	Short: "Renews an [eligible ACM certificate](https://docs.aws.amazon.com/acm/latest/userguide/managed-renewal.html).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acm_renewCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(acm_renewCertificateCmd).Standalone()

		acm_renewCertificateCmd.Flags().String("certificate-arn", "", "String that contains the ARN of the ACM certificate to be renewed.")
		acm_renewCertificateCmd.MarkFlagRequired("certificate-arn")
	})
	acmCmd.AddCommand(acm_renewCertificateCmd)
}
