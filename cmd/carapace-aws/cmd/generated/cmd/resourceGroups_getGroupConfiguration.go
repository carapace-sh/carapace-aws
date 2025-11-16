package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_getGroupConfigurationCmd = &cobra.Command{
	Use:   "get-group-configuration",
	Short: "Retrieves the service configuration associated with the specified resource group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_getGroupConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceGroups_getGroupConfigurationCmd).Standalone()

		resourceGroups_getGroupConfigurationCmd.Flags().String("group", "", "The name or the Amazon resource name (ARN) of the resource group for which you want to retrive the service configuration.")
	})
	resourceGroupsCmd.AddCommand(resourceGroups_getGroupConfigurationCmd)
}
