package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_listAssociatedGroupsCmd = &cobra.Command{
	Use:   "list-associated-groups",
	Short: "Returns a list of the groups that the specified canary is associated with.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_listAssociatedGroupsCmd).Standalone()

	synthetics_listAssociatedGroupsCmd.Flags().String("max-results", "", "Specify this parameter to limit how many groups are returned each time you use the `ListAssociatedGroups` operation.")
	synthetics_listAssociatedGroupsCmd.Flags().String("next-token", "", "A token that indicates that there is more data available.")
	synthetics_listAssociatedGroupsCmd.Flags().String("resource-arn", "", "The ARN of the canary that you want to view groups for.")
	synthetics_listAssociatedGroupsCmd.MarkFlagRequired("resource-arn")
	syntheticsCmd.AddCommand(synthetics_listAssociatedGroupsCmd)
}
