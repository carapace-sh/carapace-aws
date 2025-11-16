package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_modifyUserGroupCmd = &cobra.Command{
	Use:   "modify-user-group",
	Short: "Changes the list of users that belong to the user group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_modifyUserGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_modifyUserGroupCmd).Standalone()

		elasticache_modifyUserGroupCmd.Flags().String("engine", "", "Modifies the engine listed in a user group.")
		elasticache_modifyUserGroupCmd.Flags().String("user-group-id", "", "The ID of the user group.")
		elasticache_modifyUserGroupCmd.Flags().String("user-ids-to-add", "", "The list of user IDs to add to the user group.")
		elasticache_modifyUserGroupCmd.Flags().String("user-ids-to-remove", "", "The list of user IDs to remove from the user group.")
		elasticache_modifyUserGroupCmd.MarkFlagRequired("user-group-id")
	})
	elasticacheCmd.AddCommand(elasticache_modifyUserGroupCmd)
}
