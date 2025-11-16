package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_deletePolicyStoreCmd = &cobra.Command{
	Use:   "delete-policy-store",
	Short: "Deletes the specified policy store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_deletePolicyStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(verifiedpermissions_deletePolicyStoreCmd).Standalone()

		verifiedpermissions_deletePolicyStoreCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store that you want to delete.")
		verifiedpermissions_deletePolicyStoreCmd.MarkFlagRequired("policy-store-id")
	})
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_deletePolicyStoreCmd)
}
