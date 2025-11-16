package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_deleteUserGroupCmd = &cobra.Command{
	Use:   "delete-user-group",
	Short: "For Valkey engine version 7.2 onwards and Redis OSS 6.0 onwards: Deletes a user group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_deleteUserGroupCmd).Standalone()

	elasticache_deleteUserGroupCmd.Flags().String("user-group-id", "", "The ID of the user group.")
	elasticache_deleteUserGroupCmd.MarkFlagRequired("user-group-id")
	elasticacheCmd.AddCommand(elasticache_deleteUserGroupCmd)
}
