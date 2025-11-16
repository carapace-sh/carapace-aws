package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_describeProvisioningTemplateCmd = &cobra.Command{
	Use:   "describe-provisioning-template",
	Short: "Returns information about a provisioning template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_describeProvisioningTemplateCmd).Standalone()

	iot_describeProvisioningTemplateCmd.Flags().String("template-name", "", "The name of the provisioning template.")
	iot_describeProvisioningTemplateCmd.MarkFlagRequired("template-name")
	iotCmd.AddCommand(iot_describeProvisioningTemplateCmd)
}
