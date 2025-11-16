package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_updatePolicyStoreCmd = &cobra.Command{
	Use:   "update-policy-store",
	Short: "Modifies the validation setting for a policy store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_updatePolicyStoreCmd).Standalone()

	verifiedpermissions_updatePolicyStoreCmd.Flags().String("deletion-protection", "", "Specifies whether the policy store can be deleted.")
	verifiedpermissions_updatePolicyStoreCmd.Flags().String("description", "", "Descriptive text that you can provide to help with identification of the current policy store.")
	verifiedpermissions_updatePolicyStoreCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store that you want to update")
	verifiedpermissions_updatePolicyStoreCmd.Flags().String("validation-settings", "", "A structure that defines the validation settings that want to enable for the policy store.")
	verifiedpermissions_updatePolicyStoreCmd.MarkFlagRequired("policy-store-id")
	verifiedpermissions_updatePolicyStoreCmd.MarkFlagRequired("validation-settings")
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_updatePolicyStoreCmd)
}
