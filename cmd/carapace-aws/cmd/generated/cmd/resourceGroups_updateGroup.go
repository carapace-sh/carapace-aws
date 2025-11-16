package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_updateGroupCmd = &cobra.Command{
	Use:   "update-group",
	Short: "Updates the description for an existing group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_updateGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceGroups_updateGroupCmd).Standalone()

		resourceGroups_updateGroupCmd.Flags().String("criticality", "", "The critical rank of the application group on a scale of 1 to 10, with a rank of 1 being the most critical, and a rank of 10 being least critical.")
		resourceGroups_updateGroupCmd.Flags().String("description", "", "The new description that you want to update the resource group with.")
		resourceGroups_updateGroupCmd.Flags().String("display-name", "", "The name of the application group, which you can change at any time.")
		resourceGroups_updateGroupCmd.Flags().String("group", "", "The name or the ARN of the resource group to update.")
		resourceGroups_updateGroupCmd.Flags().String("group-name", "", "Don't use this parameter.")
		resourceGroups_updateGroupCmd.Flags().String("owner", "", "A name, email address or other identifier for the person or group who is considered as the owner of this application group within your organization.")
	})
	resourceGroupsCmd.AddCommand(resourceGroups_updateGroupCmd)
}
