package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_deleteGroupCmd = &cobra.Command{
	Use:   "delete-group",
	Short: "Deletes the specified resource group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_deleteGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceGroups_deleteGroupCmd).Standalone()

		resourceGroups_deleteGroupCmd.Flags().String("group", "", "The name or the Amazon resource name (ARN) of the resource group to delete.")
		resourceGroups_deleteGroupCmd.Flags().String("group-name", "", "Deprecated - don't use this parameter.")
	})
	resourceGroupsCmd.AddCommand(resourceGroups_deleteGroupCmd)
}
