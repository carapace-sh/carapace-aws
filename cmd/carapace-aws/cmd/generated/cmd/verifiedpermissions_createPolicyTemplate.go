package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_createPolicyTemplateCmd = &cobra.Command{
	Use:   "create-policy-template",
	Short: "Creates a policy template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_createPolicyTemplateCmd).Standalone()

	verifiedpermissions_createPolicyTemplateCmd.Flags().String("client-token", "", "Specifies a unique, case-sensitive ID that you provide to ensure the idempotency of the request.")
	verifiedpermissions_createPolicyTemplateCmd.Flags().String("description", "", "Specifies a description for the policy template.")
	verifiedpermissions_createPolicyTemplateCmd.Flags().String("policy-store-id", "", "The ID of the policy store in which to create the policy template.")
	verifiedpermissions_createPolicyTemplateCmd.Flags().String("statement", "", "Specifies the content that you want to use for the new policy template, written in the Cedar policy language.")
	verifiedpermissions_createPolicyTemplateCmd.MarkFlagRequired("policy-store-id")
	verifiedpermissions_createPolicyTemplateCmd.MarkFlagRequired("statement")
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_createPolicyTemplateCmd)
}
