package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_updateIdentitySourceCmd = &cobra.Command{
	Use:   "update-identity-source",
	Short: "Updates the specified identity source to use a new identity provider (IdP), or to change the mapping of identities from the IdP to a different principal entity type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_updateIdentitySourceCmd).Standalone()

	verifiedpermissions_updateIdentitySourceCmd.Flags().String("identity-source-id", "", "Specifies the ID of the identity source that you want to update.")
	verifiedpermissions_updateIdentitySourceCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store that contains the identity source that you want to update.")
	verifiedpermissions_updateIdentitySourceCmd.Flags().String("principal-entity-type", "", "Specifies the data type of principals generated for identities authenticated by the identity source.")
	verifiedpermissions_updateIdentitySourceCmd.Flags().String("update-configuration", "", "Specifies the details required to communicate with the identity provider (IdP) associated with this identity source.")
	verifiedpermissions_updateIdentitySourceCmd.MarkFlagRequired("identity-source-id")
	verifiedpermissions_updateIdentitySourceCmd.MarkFlagRequired("policy-store-id")
	verifiedpermissions_updateIdentitySourceCmd.MarkFlagRequired("update-configuration")
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_updateIdentitySourceCmd)
}
