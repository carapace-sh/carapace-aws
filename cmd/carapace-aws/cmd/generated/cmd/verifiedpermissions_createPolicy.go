package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_createPolicyCmd = &cobra.Command{
	Use:   "create-policy",
	Short: "Creates a Cedar policy and saves it in the specified policy store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_createPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(verifiedpermissions_createPolicyCmd).Standalone()

		verifiedpermissions_createPolicyCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive ID that you provide to ensure the idempotency of the request.")
		verifiedpermissions_createPolicyCmd.Flags().String("definition", "", "A structure that specifies the policy type and content to use for the new policy.")
		verifiedpermissions_createPolicyCmd.Flags().String("policy-store-id", "", "Specifies the `PolicyStoreId` of the policy store you want to store the policy in.")
		verifiedpermissions_createPolicyCmd.MarkFlagRequired("definition")
		verifiedpermissions_createPolicyCmd.MarkFlagRequired("policy-store-id")
	})
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_createPolicyCmd)
}
