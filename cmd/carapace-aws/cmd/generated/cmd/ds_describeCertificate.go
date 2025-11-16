package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_describeCertificateCmd = &cobra.Command{
	Use:   "describe-certificate",
	Short: "Displays information about the certificate registered for secure LDAP or client certificate authentication.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_describeCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_describeCertificateCmd).Standalone()

		ds_describeCertificateCmd.Flags().String("certificate-id", "", "The identifier of the certificate.")
		ds_describeCertificateCmd.Flags().String("directory-id", "", "The identifier of the directory.")
		ds_describeCertificateCmd.MarkFlagRequired("certificate-id")
		ds_describeCertificateCmd.MarkFlagRequired("directory-id")
	})
	dsCmd.AddCommand(ds_describeCertificateCmd)
}
