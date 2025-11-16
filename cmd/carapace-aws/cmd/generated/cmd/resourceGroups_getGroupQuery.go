package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_getGroupQueryCmd = &cobra.Command{
	Use:   "get-group-query",
	Short: "Retrieves the resource query associated with the specified resource group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_getGroupQueryCmd).Standalone()

	resourceGroups_getGroupQueryCmd.Flags().String("group", "", "The name or the Amazon resource name (ARN) of the resource group to query.")
	resourceGroups_getGroupQueryCmd.Flags().String("group-name", "", "Don't use this parameter.")
	resourceGroupsCmd.AddCommand(resourceGroups_getGroupQueryCmd)
}
