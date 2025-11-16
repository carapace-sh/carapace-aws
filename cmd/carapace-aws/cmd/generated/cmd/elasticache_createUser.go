package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "For Valkey engine version 7.2 onwards and Redis OSS 6.0 to 7.1: Creates a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_createUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_createUserCmd).Standalone()

		elasticache_createUserCmd.Flags().String("access-string", "", "Access permissions string used for this user.")
		elasticache_createUserCmd.Flags().String("authentication-mode", "", "Specifies how to authenticate the user.")
		elasticache_createUserCmd.Flags().String("engine", "", "The options are valkey or redis.")
		elasticache_createUserCmd.Flags().String("no-password-required", "", "Indicates a password is not required for this user.")
		elasticache_createUserCmd.Flags().String("passwords", "", "Passwords used for this user.")
		elasticache_createUserCmd.Flags().String("tags", "", "A list of tags to be added to this resource.")
		elasticache_createUserCmd.Flags().String("user-id", "", "The ID of the user.")
		elasticache_createUserCmd.Flags().String("user-name", "", "The username of the user.")
		elasticache_createUserCmd.MarkFlagRequired("access-string")
		elasticache_createUserCmd.MarkFlagRequired("engine")
		elasticache_createUserCmd.MarkFlagRequired("user-id")
		elasticache_createUserCmd.MarkFlagRequired("user-name")
	})
	elasticacheCmd.AddCommand(elasticache_createUserCmd)
}
