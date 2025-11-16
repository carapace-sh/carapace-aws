package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_updateProvisioningTemplateCmd = &cobra.Command{
	Use:   "update-provisioning-template",
	Short: "Updates a provisioning template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_updateProvisioningTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_updateProvisioningTemplateCmd).Standalone()

		iot_updateProvisioningTemplateCmd.Flags().String("default-version-id", "", "The ID of the default provisioning template version.")
		iot_updateProvisioningTemplateCmd.Flags().String("description", "", "The description of the provisioning template.")
		iot_updateProvisioningTemplateCmd.Flags().String("enabled", "", "True to enable the provisioning template, otherwise false.")
		iot_updateProvisioningTemplateCmd.Flags().String("pre-provisioning-hook", "", "Updates the pre-provisioning hook template.")
		iot_updateProvisioningTemplateCmd.Flags().String("provisioning-role-arn", "", "The ARN of the role associated with the provisioning template.")
		iot_updateProvisioningTemplateCmd.Flags().String("remove-pre-provisioning-hook", "", "Removes pre-provisioning hook template.")
		iot_updateProvisioningTemplateCmd.Flags().String("template-name", "", "The name of the provisioning template.")
		iot_updateProvisioningTemplateCmd.MarkFlagRequired("template-name")
	})
	iotCmd.AddCommand(iot_updateProvisioningTemplateCmd)
}
