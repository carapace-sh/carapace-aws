package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteProvisioningTemplateVersionCmd = &cobra.Command{
	Use:   "delete-provisioning-template-version",
	Short: "Deletes a provisioning template version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteProvisioningTemplateVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_deleteProvisioningTemplateVersionCmd).Standalone()

		iot_deleteProvisioningTemplateVersionCmd.Flags().String("template-name", "", "The name of the provisioning template version to delete.")
		iot_deleteProvisioningTemplateVersionCmd.Flags().String("version-id", "", "The provisioning template version ID to delete.")
		iot_deleteProvisioningTemplateVersionCmd.MarkFlagRequired("template-name")
		iot_deleteProvisioningTemplateVersionCmd.MarkFlagRequired("version-id")
	})
	iotCmd.AddCommand(iot_deleteProvisioningTemplateVersionCmd)
}
