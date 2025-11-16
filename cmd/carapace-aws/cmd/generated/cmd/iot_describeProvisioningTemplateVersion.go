package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeProvisioningTemplateVersionCmd = &cobra.Command{
	Use:   "describe-provisioning-template-version",
	Short: "Returns information about a provisioning template version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeProvisioningTemplateVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_describeProvisioningTemplateVersionCmd).Standalone()

		iot_describeProvisioningTemplateVersionCmd.Flags().String("template-name", "", "The template name.")
		iot_describeProvisioningTemplateVersionCmd.Flags().String("version-id", "", "The provisioning template version ID.")
		iot_describeProvisioningTemplateVersionCmd.MarkFlagRequired("template-name")
		iot_describeProvisioningTemplateVersionCmd.MarkFlagRequired("version-id")
	})
	iotCmd.AddCommand(iot_describeProvisioningTemplateVersionCmd)
}
