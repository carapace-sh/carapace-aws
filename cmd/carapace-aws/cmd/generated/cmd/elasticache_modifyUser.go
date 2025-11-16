package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_modifyUserCmd = &cobra.Command{
	Use:   "modify-user",
	Short: "Changes user password(s) and/or access string.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_modifyUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_modifyUserCmd).Standalone()

		elasticache_modifyUserCmd.Flags().String("access-string", "", "Access permissions string used for this user.")
		elasticache_modifyUserCmd.Flags().String("append-access-string", "", "Adds additional user permissions to the access string.")
		elasticache_modifyUserCmd.Flags().String("authentication-mode", "", "Specifies how to authenticate the user.")
		elasticache_modifyUserCmd.Flags().String("engine", "", "Modifies the engine listed for a user.")
		elasticache_modifyUserCmd.Flags().String("no-password-required", "", "Indicates no password is required for the user.")
		elasticache_modifyUserCmd.Flags().String("passwords", "", "The passwords belonging to the user.")
		elasticache_modifyUserCmd.Flags().String("user-id", "", "The ID of the user.")
		elasticache_modifyUserCmd.MarkFlagRequired("user-id")
	})
	elasticacheCmd.AddCommand(elasticache_modifyUserCmd)
}
