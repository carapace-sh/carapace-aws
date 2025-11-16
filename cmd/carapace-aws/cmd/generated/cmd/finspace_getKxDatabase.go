package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_getKxDatabaseCmd = &cobra.Command{
	Use:   "get-kx-database",
	Short: "Returns database information for the specified environment ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_getKxDatabaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspace_getKxDatabaseCmd).Standalone()

		finspace_getKxDatabaseCmd.Flags().String("database-name", "", "The name of the kdb database.")
		finspace_getKxDatabaseCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
		finspace_getKxDatabaseCmd.MarkFlagRequired("database-name")
		finspace_getKxDatabaseCmd.MarkFlagRequired("environment-id")
	})
	finspaceCmd.AddCommand(finspace_getKxDatabaseCmd)
}
