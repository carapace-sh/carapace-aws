package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_deleteKxDatabaseCmd = &cobra.Command{
	Use:   "delete-kx-database",
	Short: "Deletes the specified database and all of its associated data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_deleteKxDatabaseCmd).Standalone()

	finspace_deleteKxDatabaseCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspace_deleteKxDatabaseCmd.Flags().String("database-name", "", "The name of the kdb database that you want to delete.")
	finspace_deleteKxDatabaseCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
	finspace_deleteKxDatabaseCmd.MarkFlagRequired("client-token")
	finspace_deleteKxDatabaseCmd.MarkFlagRequired("database-name")
	finspace_deleteKxDatabaseCmd.MarkFlagRequired("environment-id")
	finspaceCmd.AddCommand(finspace_deleteKxDatabaseCmd)
}
