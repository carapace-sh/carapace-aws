package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_deleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: "For Valkey engine version 7.2 onwards and Redis OSS 6.0 onwards: Deletes a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_deleteUserCmd).Standalone()

	elasticache_deleteUserCmd.Flags().String("user-id", "", "The ID of the user.")
	elasticache_deleteUserCmd.MarkFlagRequired("user-id")
	elasticacheCmd.AddCommand(elasticache_deleteUserCmd)
}
