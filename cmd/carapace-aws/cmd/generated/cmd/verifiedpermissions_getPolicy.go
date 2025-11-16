package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_getPolicyCmd = &cobra.Command{
	Use:   "get-policy",
	Short: "Retrieves information about the specified policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_getPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(verifiedpermissions_getPolicyCmd).Standalone()

		verifiedpermissions_getPolicyCmd.Flags().String("policy-id", "", "Specifies the ID of the policy you want information about.")
		verifiedpermissions_getPolicyCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store that contains the policy that you want information about.")
		verifiedpermissions_getPolicyCmd.MarkFlagRequired("policy-id")
		verifiedpermissions_getPolicyCmd.MarkFlagRequired("policy-store-id")
	})
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_getPolicyCmd)
}
