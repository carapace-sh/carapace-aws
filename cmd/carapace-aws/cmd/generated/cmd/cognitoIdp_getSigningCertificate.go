package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_getSigningCertificateCmd = &cobra.Command{
	Use:   "get-signing-certificate",
	Short: "Given a user pool ID, returns the signing certificate for SAML 2.0 federation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_getSigningCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_getSigningCertificateCmd).Standalone()

		cognitoIdp_getSigningCertificateCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to view the signing certificate.")
		cognitoIdp_getSigningCertificateCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_getSigningCertificateCmd)
}
