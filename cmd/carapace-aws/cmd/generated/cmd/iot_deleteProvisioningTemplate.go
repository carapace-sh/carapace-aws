package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_deleteProvisioningTemplateCmd = &cobra.Command{
	Use:   "delete-provisioning-template",
	Short: "Deletes a provisioning template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_deleteProvisioningTemplateCmd).Standalone()

	iot_deleteProvisioningTemplateCmd.Flags().String("template-name", "", "The name of the fleet provision template to delete.")
	iot_deleteProvisioningTemplateCmd.MarkFlagRequired("template-name")
	iotCmd.AddCommand(iot_deleteProvisioningTemplateCmd)
}
