package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_putGroupConfigurationCmd = &cobra.Command{
	Use:   "put-group-configuration",
	Short: "Attaches a service configuration to the specified group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_putGroupConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceGroups_putGroupConfigurationCmd).Standalone()

		resourceGroups_putGroupConfigurationCmd.Flags().String("configuration", "", "The new configuration to associate with the specified group.")
		resourceGroups_putGroupConfigurationCmd.Flags().String("group", "", "The name or Amazon resource name (ARN) of the resource group with the configuration that you want to update.")
	})
	resourceGroupsCmd.AddCommand(resourceGroups_putGroupConfigurationCmd)
}
