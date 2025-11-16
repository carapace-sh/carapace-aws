package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_createProvisioningTemplateCmd = &cobra.Command{
	Use:   "create-provisioning-template",
	Short: "Creates a provisioning template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_createProvisioningTemplateCmd).Standalone()

	iot_createProvisioningTemplateCmd.Flags().String("description", "", "The description of the provisioning template.")
	iot_createProvisioningTemplateCmd.Flags().String("enabled", "", "True to enable the provisioning template, otherwise false.")
	iot_createProvisioningTemplateCmd.Flags().String("pre-provisioning-hook", "", "Creates a pre-provisioning hook template.")
	iot_createProvisioningTemplateCmd.Flags().String("provisioning-role-arn", "", "The role ARN for the role associated with the provisioning template.")
	iot_createProvisioningTemplateCmd.Flags().String("tags", "", "Metadata which can be used to manage the provisioning template.")
	iot_createProvisioningTemplateCmd.Flags().String("template-body", "", "The JSON formatted contents of the provisioning template.")
	iot_createProvisioningTemplateCmd.Flags().String("template-name", "", "The name of the provisioning template.")
	iot_createProvisioningTemplateCmd.Flags().String("type", "", "The type you define in a provisioning template.")
	iot_createProvisioningTemplateCmd.MarkFlagRequired("provisioning-role-arn")
	iot_createProvisioningTemplateCmd.MarkFlagRequired("template-body")
	iot_createProvisioningTemplateCmd.MarkFlagRequired("template-name")
	iotCmd.AddCommand(iot_createProvisioningTemplateCmd)
}
