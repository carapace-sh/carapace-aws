package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listProvisioningTemplateVersionsCmd = &cobra.Command{
	Use:   "list-provisioning-template-versions",
	Short: "A list of provisioning template versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listProvisioningTemplateVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_listProvisioningTemplateVersionsCmd).Standalone()

		iot_listProvisioningTemplateVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
		iot_listProvisioningTemplateVersionsCmd.Flags().String("next-token", "", "A token to retrieve the next set of results.")
		iot_listProvisioningTemplateVersionsCmd.Flags().String("template-name", "", "The name of the provisioning template.")
		iot_listProvisioningTemplateVersionsCmd.MarkFlagRequired("template-name")
	})
	iotCmd.AddCommand(iot_listProvisioningTemplateVersionsCmd)
}
