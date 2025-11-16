package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_updateGroupQueryCmd = &cobra.Command{
	Use:   "update-group-query",
	Short: "Updates the resource query of a group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_updateGroupQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceGroups_updateGroupQueryCmd).Standalone()

		resourceGroups_updateGroupQueryCmd.Flags().String("group", "", "The name or the Amazon resource name (ARN) of the resource group to query.")
		resourceGroups_updateGroupQueryCmd.Flags().String("group-name", "", "Don't use this parameter.")
		resourceGroups_updateGroupQueryCmd.Flags().String("resource-query", "", "The resource query to determine which Amazon Web Services resources are members of this resource group.")
		resourceGroups_updateGroupQueryCmd.MarkFlagRequired("resource-query")
	})
	resourceGroupsCmd.AddCommand(resourceGroups_updateGroupQueryCmd)
}
