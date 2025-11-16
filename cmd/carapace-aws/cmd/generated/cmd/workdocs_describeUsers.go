package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_describeUsersCmd = &cobra.Command{
	Use:   "describe-users",
	Short: "Describes the specified users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_describeUsersCmd).Standalone()

	workdocs_describeUsersCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
	workdocs_describeUsersCmd.Flags().String("fields", "", "A comma-separated list of values.")
	workdocs_describeUsersCmd.Flags().String("include", "", "The state of the users.")
	workdocs_describeUsersCmd.Flags().String("limit", "", "The maximum number of items to return.")
	workdocs_describeUsersCmd.Flags().String("marker", "", "The marker for the next set of results.")
	workdocs_describeUsersCmd.Flags().String("order", "", "The order for the results.")
	workdocs_describeUsersCmd.Flags().String("organization-id", "", "The ID of the organization.")
	workdocs_describeUsersCmd.Flags().String("query", "", "A query to filter users by user name.")
	workdocs_describeUsersCmd.Flags().String("sort", "", "The sorting criteria.")
	workdocs_describeUsersCmd.Flags().String("user-ids", "", "The IDs of the users.")
	workdocsCmd.AddCommand(workdocs_describeUsersCmd)
}
