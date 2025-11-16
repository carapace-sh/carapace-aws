package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: "Provides summary information about the users for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listUsersCmd).Standalone()

	connect_listUsersCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_listUsersCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_listUsersCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_listUsersCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_listUsersCmd)
}
