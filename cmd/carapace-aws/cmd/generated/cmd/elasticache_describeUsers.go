package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_describeUsersCmd = &cobra.Command{
	Use:   "describe-users",
	Short: "Returns a list of users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_describeUsersCmd).Standalone()

	elasticache_describeUsersCmd.Flags().String("engine", "", "The engine.")
	elasticache_describeUsersCmd.Flags().String("filters", "", "Filter to determine the list of User IDs to return.")
	elasticache_describeUsersCmd.Flags().String("marker", "", "An optional marker returned from a prior request.")
	elasticache_describeUsersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	elasticache_describeUsersCmd.Flags().String("user-id", "", "The ID of the user.")
	elasticacheCmd.AddCommand(elasticache_describeUsersCmd)
}
