package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_deregisterCertificateCmd = &cobra.Command{
	Use:   "deregister-certificate",
	Short: "Deletes from the system the certificate that was registered for secure LDAP or client certificate authentication.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_deregisterCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_deregisterCertificateCmd).Standalone()

		ds_deregisterCertificateCmd.Flags().String("certificate-id", "", "The identifier of the certificate.")
		ds_deregisterCertificateCmd.Flags().String("directory-id", "", "The identifier of the directory.")
		ds_deregisterCertificateCmd.MarkFlagRequired("certificate-id")
		ds_deregisterCertificateCmd.MarkFlagRequired("directory-id")
	})
	dsCmd.AddCommand(ds_deregisterCertificateCmd)
}
