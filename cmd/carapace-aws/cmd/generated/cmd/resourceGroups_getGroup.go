package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_getGroupCmd = &cobra.Command{
	Use:   "get-group",
	Short: "Returns information about a specified resource group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_getGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceGroups_getGroupCmd).Standalone()

		resourceGroups_getGroupCmd.Flags().String("group", "", "The name or the Amazon resource name (ARN) of the resource group to retrieve.")
		resourceGroups_getGroupCmd.Flags().String("group-name", "", "Deprecated - don't use this parameter.")
	})
	resourceGroupsCmd.AddCommand(resourceGroups_getGroupCmd)
}
