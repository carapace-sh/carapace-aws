package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acm_describeCertificateCmd = &cobra.Command{
	Use:   "describe-certificate",
	Short: "Returns detailed metadata about the specified ACM certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acm_describeCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(acm_describeCertificateCmd).Standalone()

		acm_describeCertificateCmd.Flags().String("certificate-arn", "", "The Amazon Resource Name (ARN) of the ACM certificate.")
		acm_describeCertificateCmd.MarkFlagRequired("certificate-arn")
	})
	acmCmd.AddCommand(acm_describeCertificateCmd)
}
