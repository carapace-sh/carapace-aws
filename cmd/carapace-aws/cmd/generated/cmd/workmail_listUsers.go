package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_listUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: "Returns summaries of the organization's users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_listUsersCmd).Standalone()

	workmail_listUsersCmd.Flags().String("filters", "", "Limit the user search results based on the filter criteria.")
	workmail_listUsersCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	workmail_listUsersCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
	workmail_listUsersCmd.Flags().String("organization-id", "", "The identifier for the organization under which the users exist.")
	workmail_listUsersCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_listUsersCmd)
}
