package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getRelationalDatabaseCmd = &cobra.Command{
	Use:   "get-relational-database",
	Short: "Returns information about a specific database in Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getRelationalDatabaseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getRelationalDatabaseCmd).Standalone()

		lightsail_getRelationalDatabaseCmd.Flags().String("relational-database-name", "", "The name of the database that you are looking up.")
		lightsail_getRelationalDatabaseCmd.MarkFlagRequired("relational-database-name")
	})
	lightsailCmd.AddCommand(lightsail_getRelationalDatabaseCmd)
}
