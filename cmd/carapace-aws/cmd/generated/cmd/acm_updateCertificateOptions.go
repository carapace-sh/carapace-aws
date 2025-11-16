package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acm_updateCertificateOptionsCmd = &cobra.Command{
	Use:   "update-certificate-options",
	Short: "Updates a certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acm_updateCertificateOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(acm_updateCertificateOptionsCmd).Standalone()

		acm_updateCertificateOptionsCmd.Flags().String("certificate-arn", "", "ARN of the requested certificate to update.")
		acm_updateCertificateOptionsCmd.Flags().String("options", "", "Use to update the options for your certificate.")
		acm_updateCertificateOptionsCmd.MarkFlagRequired("certificate-arn")
		acm_updateCertificateOptionsCmd.MarkFlagRequired("options")
	})
	acmCmd.AddCommand(acm_updateCertificateOptionsCmd)
}
