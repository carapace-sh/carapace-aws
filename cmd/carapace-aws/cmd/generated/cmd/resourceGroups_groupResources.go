package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_groupResourcesCmd = &cobra.Command{
	Use:   "group-resources",
	Short: "Adds the specified resources to the specified group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_groupResourcesCmd).Standalone()

	resourceGroups_groupResourcesCmd.Flags().String("group", "", "The name or the Amazon resource name (ARN) of the resource group to add resources to.")
	resourceGroups_groupResourcesCmd.Flags().String("resource-arns", "", "The list of Amazon resource names (ARNs) of the resources to be added to the group.")
	resourceGroups_groupResourcesCmd.MarkFlagRequired("group")
	resourceGroups_groupResourcesCmd.MarkFlagRequired("resource-arns")
	resourceGroupsCmd.AddCommand(resourceGroups_groupResourcesCmd)
}
