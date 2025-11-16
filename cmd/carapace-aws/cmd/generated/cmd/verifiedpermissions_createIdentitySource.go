package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_createIdentitySourceCmd = &cobra.Command{
	Use:   "create-identity-source",
	Short: "Adds an identity source to a policy store–an Amazon Cognito user pool or OpenID Connect (OIDC) identity provider (IdP).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_createIdentitySourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(verifiedpermissions_createIdentitySourceCmd).Standalone()

		verifiedpermissions_createIdentitySourceCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive ID that you provide to ensure the idempotency of the request.")
		verifiedpermissions_createIdentitySourceCmd.Flags().String("configuration", "", "Specifies the details required to communicate with the identity provider (IdP) associated with this identity source.")
		verifiedpermissions_createIdentitySourceCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store in which you want to store this identity source.")
		verifiedpermissions_createIdentitySourceCmd.Flags().String("principal-entity-type", "", "Specifies the namespace and data type of the principals generated for identities authenticated by the new identity source.")
		verifiedpermissions_createIdentitySourceCmd.MarkFlagRequired("configuration")
		verifiedpermissions_createIdentitySourceCmd.MarkFlagRequired("policy-store-id")
	})
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_createIdentitySourceCmd)
}
