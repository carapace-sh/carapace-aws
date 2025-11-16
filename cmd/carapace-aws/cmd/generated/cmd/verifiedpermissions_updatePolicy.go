package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_updatePolicyCmd = &cobra.Command{
	Use:   "update-policy",
	Short: "Modifies a Cedar static policy in the specified policy store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_updatePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(verifiedpermissions_updatePolicyCmd).Standalone()

		verifiedpermissions_updatePolicyCmd.Flags().String("definition", "", "Specifies the updated policy content that you want to replace on the specified policy.")
		verifiedpermissions_updatePolicyCmd.Flags().String("policy-id", "", "Specifies the ID of the policy that you want to update.")
		verifiedpermissions_updatePolicyCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store that contains the policy that you want to update.")
		verifiedpermissions_updatePolicyCmd.MarkFlagRequired("definition")
		verifiedpermissions_updatePolicyCmd.MarkFlagRequired("policy-id")
		verifiedpermissions_updatePolicyCmd.MarkFlagRequired("policy-store-id")
	})
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_updatePolicyCmd)
}
