package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_createUserGroupCmd = &cobra.Command{
	Use:   "create-user-group",
	Short: "For Valkey engine version 7.2 onwards and Redis OSS 6.0 to 7.1: Creates a user group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_createUserGroupCmd).Standalone()

	elasticache_createUserGroupCmd.Flags().String("engine", "", "Sets the engine listed in a user group.")
	elasticache_createUserGroupCmd.Flags().String("tags", "", "A list of tags to be added to this resource.")
	elasticache_createUserGroupCmd.Flags().String("user-group-id", "", "The ID of the user group.")
	elasticache_createUserGroupCmd.Flags().String("user-ids", "", "The list of user IDs that belong to the user group.")
	elasticache_createUserGroupCmd.MarkFlagRequired("engine")
	elasticache_createUserGroupCmd.MarkFlagRequired("user-group-id")
	elasticacheCmd.AddCommand(elasticache_createUserGroupCmd)
}
