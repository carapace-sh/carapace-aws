package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_ungroupResourcesCmd = &cobra.Command{
	Use:   "ungroup-resources",
	Short: "Removes the specified resources from the specified group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_ungroupResourcesCmd).Standalone()

	resourceGroups_ungroupResourcesCmd.Flags().String("group", "", "The name or the Amazon resource name (ARN) of the resource group from which to remove the resources.")
	resourceGroups_ungroupResourcesCmd.Flags().String("resource-arns", "", "The Amazon resource names (ARNs) of the resources to be removed from the group.")
	resourceGroups_ungroupResourcesCmd.MarkFlagRequired("group")
	resourceGroups_ungroupResourcesCmd.MarkFlagRequired("resource-arns")
	resourceGroupsCmd.AddCommand(resourceGroups_ungroupResourcesCmd)
}
