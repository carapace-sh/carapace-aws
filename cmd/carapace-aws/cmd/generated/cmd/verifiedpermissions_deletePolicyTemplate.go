package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_deletePolicyTemplateCmd = &cobra.Command{
	Use:   "delete-policy-template",
	Short: "Deletes the specified policy template from the policy store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_deletePolicyTemplateCmd).Standalone()

	verifiedpermissions_deletePolicyTemplateCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store that contains the policy template that you want to delete.")
	verifiedpermissions_deletePolicyTemplateCmd.Flags().String("policy-template-id", "", "Specifies the ID of the policy template that you want to delete.")
	verifiedpermissions_deletePolicyTemplateCmd.MarkFlagRequired("policy-store-id")
	verifiedpermissions_deletePolicyTemplateCmd.MarkFlagRequired("policy-template-id")
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_deletePolicyTemplateCmd)
}
