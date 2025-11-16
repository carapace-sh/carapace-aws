package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acm_deleteCertificateCmd = &cobra.Command{
	Use:   "delete-certificate",
	Short: "Deletes a certificate and its associated private key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acm_deleteCertificateCmd).Standalone()

	acm_deleteCertificateCmd.Flags().String("certificate-arn", "", "String that contains the ARN of the ACM certificate to be deleted.")
	acm_deleteCertificateCmd.MarkFlagRequired("certificate-arn")
	acmCmd.AddCommand(acm_deleteCertificateCmd)
}
