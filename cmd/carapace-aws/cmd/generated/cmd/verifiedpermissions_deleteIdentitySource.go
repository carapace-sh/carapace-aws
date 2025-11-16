package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_deleteIdentitySourceCmd = &cobra.Command{
	Use:   "delete-identity-source",
	Short: "Deletes an identity source that references an identity provider (IdP) such as Amazon Cognito.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_deleteIdentitySourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(verifiedpermissions_deleteIdentitySourceCmd).Standalone()

		verifiedpermissions_deleteIdentitySourceCmd.Flags().String("identity-source-id", "", "Specifies the ID of the identity source that you want to delete.")
		verifiedpermissions_deleteIdentitySourceCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store that contains the identity source that you want to delete.")
		verifiedpermissions_deleteIdentitySourceCmd.MarkFlagRequired("identity-source-id")
		verifiedpermissions_deleteIdentitySourceCmd.MarkFlagRequired("policy-store-id")
	})
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_deleteIdentitySourceCmd)
}
