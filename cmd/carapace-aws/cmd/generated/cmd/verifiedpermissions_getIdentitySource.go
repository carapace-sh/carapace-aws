package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_getIdentitySourceCmd = &cobra.Command{
	Use:   "get-identity-source",
	Short: "Retrieves the details about the specified identity source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_getIdentitySourceCmd).Standalone()

	verifiedpermissions_getIdentitySourceCmd.Flags().String("identity-source-id", "", "Specifies the ID of the identity source you want information about.")
	verifiedpermissions_getIdentitySourceCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store that contains the identity source you want information about.")
	verifiedpermissions_getIdentitySourceCmd.MarkFlagRequired("identity-source-id")
	verifiedpermissions_getIdentitySourceCmd.MarkFlagRequired("policy-store-id")
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_getIdentitySourceCmd)
}
