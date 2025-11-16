package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_deletePolicyCmd = &cobra.Command{
	Use:   "delete-policy",
	Short: "Deletes the specified policy from the policy store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_deletePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(verifiedpermissions_deletePolicyCmd).Standalone()

		verifiedpermissions_deletePolicyCmd.Flags().String("policy-id", "", "Specifies the ID of the policy that you want to delete.")
		verifiedpermissions_deletePolicyCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store that contains the policy that you want to delete.")
		verifiedpermissions_deletePolicyCmd.MarkFlagRequired("policy-id")
		verifiedpermissions_deletePolicyCmd.MarkFlagRequired("policy-store-id")
	})
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_deletePolicyCmd)
}
