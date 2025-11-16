package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_describeUserGroupsCmd = &cobra.Command{
	Use:   "describe-user-groups",
	Short: "Returns a list of user groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_describeUserGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_describeUserGroupsCmd).Standalone()

		elasticache_describeUserGroupsCmd.Flags().String("marker", "", "An optional marker returned from a prior request.")
		elasticache_describeUserGroupsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		elasticache_describeUserGroupsCmd.Flags().String("user-group-id", "", "The ID of the user group.")
	})
	elasticacheCmd.AddCommand(elasticache_describeUserGroupsCmd)
}
