package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_listPolicyTemplatesCmd = &cobra.Command{
	Use:   "list-policy-templates",
	Short: "Returns a paginated list of all policy templates in the specified policy store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_listPolicyTemplatesCmd).Standalone()

	verifiedpermissions_listPolicyTemplatesCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included in each response.")
	verifiedpermissions_listPolicyTemplatesCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
	verifiedpermissions_listPolicyTemplatesCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store that contains the policy templates you want to list.")
	verifiedpermissions_listPolicyTemplatesCmd.MarkFlagRequired("policy-store-id")
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_listPolicyTemplatesCmd)
}
