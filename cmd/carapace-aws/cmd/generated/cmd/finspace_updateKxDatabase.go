package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_updateKxDatabaseCmd = &cobra.Command{
	Use:   "update-kx-database",
	Short: "Updates information for the given kdb database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_updateKxDatabaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspace_updateKxDatabaseCmd).Standalone()

		finspace_updateKxDatabaseCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
		finspace_updateKxDatabaseCmd.Flags().String("database-name", "", "The name of the kdb database.")
		finspace_updateKxDatabaseCmd.Flags().String("description", "", "A description of the database.")
		finspace_updateKxDatabaseCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
		finspace_updateKxDatabaseCmd.MarkFlagRequired("client-token")
		finspace_updateKxDatabaseCmd.MarkFlagRequired("database-name")
		finspace_updateKxDatabaseCmd.MarkFlagRequired("environment-id")
	})
	finspaceCmd.AddCommand(finspace_updateKxDatabaseCmd)
}
