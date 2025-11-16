package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_getPolicyTemplateCmd = &cobra.Command{
	Use:   "get-policy-template",
	Short: "Retrieve the details for the specified policy template in the specified policy store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_getPolicyTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(verifiedpermissions_getPolicyTemplateCmd).Standalone()

		verifiedpermissions_getPolicyTemplateCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store that contains the policy template that you want information about.")
		verifiedpermissions_getPolicyTemplateCmd.Flags().String("policy-template-id", "", "Specifies the ID of the policy template that you want information about.")
		verifiedpermissions_getPolicyTemplateCmd.MarkFlagRequired("policy-store-id")
		verifiedpermissions_getPolicyTemplateCmd.MarkFlagRequired("policy-template-id")
	})
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_getPolicyTemplateCmd)
}
