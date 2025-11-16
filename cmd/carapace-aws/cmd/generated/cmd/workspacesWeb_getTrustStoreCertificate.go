package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesWeb_getTrustStoreCertificateCmd = &cobra.Command{
	Use:   "get-trust-store-certificate",
	Short: "Gets the trust store certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesWeb_getTrustStoreCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesWeb_getTrustStoreCertificateCmd).Standalone()

		workspacesWeb_getTrustStoreCertificateCmd.Flags().String("thumbprint", "", "The thumbprint of the trust store certificate.")
		workspacesWeb_getTrustStoreCertificateCmd.Flags().String("trust-store-arn", "", "The ARN of the trust store certificate.")
		workspacesWeb_getTrustStoreCertificateCmd.MarkFlagRequired("thumbprint")
		workspacesWeb_getTrustStoreCertificateCmd.MarkFlagRequired("trust-store-arn")
	})
	workspacesWebCmd.AddCommand(workspacesWeb_getTrustStoreCertificateCmd)
}
