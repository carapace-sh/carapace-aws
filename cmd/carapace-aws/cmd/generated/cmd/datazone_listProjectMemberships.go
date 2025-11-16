package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listProjectMembershipsCmd = &cobra.Command{
	Use:   "list-project-memberships",
	Short: "Lists all members of the specified project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listProjectMembershipsCmd).Standalone()

	datazone_listProjectMembershipsCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which you want to list project memberships.")
	datazone_listProjectMembershipsCmd.Flags().String("max-results", "", "The maximum number of memberships to return in a single call to `ListProjectMemberships`.")
	datazone_listProjectMembershipsCmd.Flags().String("next-token", "", "When the number of memberships is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of memberships, the response includes a pagination token named `NextToken`.")
	datazone_listProjectMembershipsCmd.Flags().String("project-identifier", "", "The identifier of the project whose memberships you want to list.")
	datazone_listProjectMembershipsCmd.Flags().String("sort-by", "", "The method by which you want to sort the project memberships.")
	datazone_listProjectMembershipsCmd.Flags().String("sort-order", "", "The sort order of the project memberships.")
	datazone_listProjectMembershipsCmd.MarkFlagRequired("domain-identifier")
	datazone_listProjectMembershipsCmd.MarkFlagRequired("project-identifier")
	datazoneCmd.AddCommand(datazone_listProjectMembershipsCmd)
}
