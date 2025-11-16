package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createProvisioningTemplateVersionCmd = &cobra.Command{
	Use:   "create-provisioning-template-version",
	Short: "Creates a new version of a provisioning template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createProvisioningTemplateVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_createProvisioningTemplateVersionCmd).Standalone()

		iot_createProvisioningTemplateVersionCmd.Flags().String("set-as-default", "", "Sets a fleet provision template version as the default version.")
		iot_createProvisioningTemplateVersionCmd.Flags().String("template-body", "", "The JSON formatted contents of the provisioning template.")
		iot_createProvisioningTemplateVersionCmd.Flags().String("template-name", "", "The name of the provisioning template.")
		iot_createProvisioningTemplateVersionCmd.MarkFlagRequired("template-body")
		iot_createProvisioningTemplateVersionCmd.MarkFlagRequired("template-name")
	})
	iotCmd.AddCommand(iot_createProvisioningTemplateVersionCmd)
}
