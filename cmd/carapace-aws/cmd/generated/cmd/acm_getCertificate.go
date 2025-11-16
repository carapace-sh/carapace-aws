package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acm_getCertificateCmd = &cobra.Command{
	Use:   "get-certificate",
	Short: "Retrieves a certificate and its certificate chain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acm_getCertificateCmd).Standalone()

	acm_getCertificateCmd.Flags().String("certificate-arn", "", "String that contains a certificate ARN in the following format:")
	acm_getCertificateCmd.MarkFlagRequired("certificate-arn")
	acmCmd.AddCommand(acm_getCertificateCmd)
}
