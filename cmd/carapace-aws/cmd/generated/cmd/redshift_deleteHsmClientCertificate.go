package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deleteHsmClientCertificateCmd = &cobra.Command{
	Use:   "delete-hsm-client-certificate",
	Short: "Deletes the specified HSM client certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deleteHsmClientCertificateCmd).Standalone()

	redshift_deleteHsmClientCertificateCmd.Flags().String("hsm-client-certificate-identifier", "", "The identifier of the HSM client certificate to be deleted.")
	redshift_deleteHsmClientCertificateCmd.MarkFlagRequired("hsm-client-certificate-identifier")
	redshiftCmd.AddCommand(redshift_deleteHsmClientCertificateCmd)
}
