package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_registerCertificateCmd = &cobra.Command{
	Use:   "register-certificate",
	Short: "Registers a certificate for a secure LDAP or client certificate authentication.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_registerCertificateCmd).Standalone()

	ds_registerCertificateCmd.Flags().String("certificate-data", "", "The certificate PEM string that needs to be registered.")
	ds_registerCertificateCmd.Flags().String("client-cert-auth-settings", "", "A `ClientCertAuthSettings` object that contains client certificate authentication settings.")
	ds_registerCertificateCmd.Flags().String("directory-id", "", "The identifier of the directory.")
	ds_registerCertificateCmd.Flags().String("type", "", "The function that the registered certificate performs.")
	ds_registerCertificateCmd.MarkFlagRequired("certificate-data")
	ds_registerCertificateCmd.MarkFlagRequired("directory-id")
	dsCmd.AddCommand(ds_registerCertificateCmd)
}
