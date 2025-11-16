package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_createHsmClientCertificateCmd = &cobra.Command{
	Use:   "create-hsm-client-certificate",
	Short: "Creates an HSM client certificate that an Amazon Redshift cluster will use to connect to the client's HSM in order to store and retrieve the keys used to encrypt the cluster databases.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_createHsmClientCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_createHsmClientCertificateCmd).Standalone()

		redshift_createHsmClientCertificateCmd.Flags().String("hsm-client-certificate-identifier", "", "The identifier to be assigned to the new HSM client certificate that the cluster will use to connect to the HSM to use the database encryption keys.")
		redshift_createHsmClientCertificateCmd.Flags().String("tags", "", "A list of tag instances.")
		redshift_createHsmClientCertificateCmd.MarkFlagRequired("hsm-client-certificate-identifier")
	})
	redshiftCmd.AddCommand(redshift_createHsmClientCertificateCmd)
}
