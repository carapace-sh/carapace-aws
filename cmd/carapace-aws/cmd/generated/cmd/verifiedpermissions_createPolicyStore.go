package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_createPolicyStoreCmd = &cobra.Command{
	Use:   "create-policy-store",
	Short: "Creates a policy store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_createPolicyStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(verifiedpermissions_createPolicyStoreCmd).Standalone()

		verifiedpermissions_createPolicyStoreCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive ID that you provide to ensure the idempotency of the request.")
		verifiedpermissions_createPolicyStoreCmd.Flags().String("deletion-protection", "", "Specifies whether the policy store can be deleted.")
		verifiedpermissions_createPolicyStoreCmd.Flags().String("description", "", "Descriptive text that you can provide to help with identification of the current policy store.")
		verifiedpermissions_createPolicyStoreCmd.Flags().String("tags", "", "The list of key-value pairs to associate with the policy store.")
		verifiedpermissions_createPolicyStoreCmd.Flags().String("validation-settings", "", "Specifies the validation setting for this policy store.")
		verifiedpermissions_createPolicyStoreCmd.MarkFlagRequired("validation-settings")
	})
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_createPolicyStoreCmd)
}
