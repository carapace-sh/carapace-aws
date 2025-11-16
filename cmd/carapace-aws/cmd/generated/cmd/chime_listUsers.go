package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_listUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: "Lists the users that belong to the specified Amazon Chime account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_listUsersCmd).Standalone()

	chime_listUsersCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
	chime_listUsersCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	chime_listUsersCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
	chime_listUsersCmd.Flags().String("user-email", "", "Optional.")
	chime_listUsersCmd.Flags().String("user-type", "", "The user type.")
	chime_listUsersCmd.MarkFlagRequired("account-id")
	chimeCmd.AddCommand(chime_listUsersCmd)
}
