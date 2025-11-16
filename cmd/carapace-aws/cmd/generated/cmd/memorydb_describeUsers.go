package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_describeUsersCmd = &cobra.Command{
	Use:   "describe-users",
	Short: "Returns a list of users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_describeUsersCmd).Standalone()

	memorydb_describeUsersCmd.Flags().String("filters", "", "Filter to determine the list of users to return.")
	memorydb_describeUsersCmd.Flags().String("max-results", "", "The maximum number of records to include in the response.")
	memorydb_describeUsersCmd.Flags().String("next-token", "", "An optional argument to pass in case the total number of records exceeds the value of MaxResults.")
	memorydb_describeUsersCmd.Flags().String("user-name", "", "The name of the user.")
	memorydbCmd.AddCommand(memorydb_describeUsersCmd)
}
