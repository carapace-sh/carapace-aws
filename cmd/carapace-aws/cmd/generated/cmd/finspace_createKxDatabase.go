package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_createKxDatabaseCmd = &cobra.Command{
	Use:   "create-kx-database",
	Short: "Creates a new kdb database in the environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_createKxDatabaseCmd).Standalone()

	finspace_createKxDatabaseCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspace_createKxDatabaseCmd.Flags().String("database-name", "", "The name of the kdb database.")
	finspace_createKxDatabaseCmd.Flags().String("description", "", "A description of the database.")
	finspace_createKxDatabaseCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
	finspace_createKxDatabaseCmd.Flags().String("tags", "", "A list of key-value pairs to label the kdb database.")
	finspace_createKxDatabaseCmd.MarkFlagRequired("client-token")
	finspace_createKxDatabaseCmd.MarkFlagRequired("database-name")
	finspace_createKxDatabaseCmd.MarkFlagRequired("environment-id")
	finspaceCmd.AddCommand(finspace_createKxDatabaseCmd)
}
