package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_describeGroupsCmd = &cobra.Command{
	Use:   "describe-groups",
	Short: "Describes the groups specified by the query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_describeGroupsCmd).Standalone()

	workdocs_describeGroupsCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_describeGroupsCmd.Flags().String("limit", "", "The maximum number of items to return with this call.")
	workdocs_describeGroupsCmd.Flags().String("marker", "", "The marker for the next set of results.")
	workdocs_describeGroupsCmd.Flags().String("organization-id", "", "The ID of the organization.")
	workdocs_describeGroupsCmd.Flags().String("search-query", "", "A query to describe groups by group name.")
	workdocs_describeGroupsCmd.MarkFlagRequired("search-query")
	workdocsCmd.AddCommand(workdocs_describeGroupsCmd)
}
