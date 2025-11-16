package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_listProvisioningTemplatesCmd = &cobra.Command{
	Use:   "list-provisioning-templates",
	Short: "Lists the provisioning templates in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_listProvisioningTemplatesCmd).Standalone()

	iot_listProvisioningTemplatesCmd.Flags().String("max-results", "", "The maximum number of results to return at one time.")
	iot_listProvisioningTemplatesCmd.Flags().String("next-token", "", "A token to retrieve the next set of results.")
	iotCmd.AddCommand(iot_listProvisioningTemplatesCmd)
}
