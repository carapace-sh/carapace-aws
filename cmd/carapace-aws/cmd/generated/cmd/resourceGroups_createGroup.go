package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceGroups_createGroupCmd = &cobra.Command{
	Use:   "create-group",
	Short: "Creates a resource group with the specified name and description.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceGroups_createGroupCmd).Standalone()

	resourceGroups_createGroupCmd.Flags().String("configuration", "", "A configuration associates the resource group with an Amazon Web Services service and specifies how the service can interact with the resources in the group.")
	resourceGroups_createGroupCmd.Flags().String("criticality", "", "The critical rank of the application group on a scale of 1 to 10, with a rank of 1 being the most critical, and a rank of 10 being least critical.")
	resourceGroups_createGroupCmd.Flags().String("description", "", "The description of the resource group.")
	resourceGroups_createGroupCmd.Flags().String("display-name", "", "The name of the application group, which you can change at any time.")
	resourceGroups_createGroupCmd.Flags().String("name", "", "The name of the group, which is the identifier of the group in other operations.")
	resourceGroups_createGroupCmd.Flags().String("owner", "", "A name, email address or other identifier for the person or group who is considered as the owner of this application group within your organization.")
	resourceGroups_createGroupCmd.Flags().String("resource-query", "", "The resource query that determines which Amazon Web Services resources are members of this group.")
	resourceGroups_createGroupCmd.Flags().String("tags", "", "The tags to add to the group.")
	resourceGroups_createGroupCmd.MarkFlagRequired("name")
	resourceGroupsCmd.AddCommand(resourceGroups_createGroupCmd)
}
