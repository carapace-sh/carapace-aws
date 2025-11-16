package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getRelationalDatabaseLogStreamsCmd = &cobra.Command{
	Use:   "get-relational-database-log-streams",
	Short: "Returns a list of available log streams for a specific database in Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getRelationalDatabaseLogStreamsCmd).Standalone()

	lightsail_getRelationalDatabaseLogStreamsCmd.Flags().String("relational-database-name", "", "The name of your database for which to get log streams.")
	lightsail_getRelationalDatabaseLogStreamsCmd.MarkFlagRequired("relational-database-name")
	lightsailCmd.AddCommand(lightsail_getRelationalDatabaseLogStreamsCmd)
}
