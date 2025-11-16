package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_updatePolicyTemplateCmd = &cobra.Command{
	Use:   "update-policy-template",
	Short: "Updates the specified policy template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_updatePolicyTemplateCmd).Standalone()

	verifiedpermissions_updatePolicyTemplateCmd.Flags().String("description", "", "Specifies a new description to apply to the policy template.")
	verifiedpermissions_updatePolicyTemplateCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store that contains the policy template that you want to update.")
	verifiedpermissions_updatePolicyTemplateCmd.Flags().String("policy-template-id", "", "Specifies the ID of the policy template that you want to update.")
	verifiedpermissions_updatePolicyTemplateCmd.Flags().String("statement", "", "Specifies new statement content written in Cedar policy language to replace the current body of the policy template.")
	verifiedpermissions_updatePolicyTemplateCmd.MarkFlagRequired("policy-store-id")
	verifiedpermissions_updatePolicyTemplateCmd.MarkFlagRequired("policy-template-id")
	verifiedpermissions_updatePolicyTemplateCmd.MarkFlagRequired("statement")
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_updatePolicyTemplateCmd)
}
