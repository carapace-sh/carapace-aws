package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getRelationalDatabaseParametersCmd = &cobra.Command{
	Use:   "get-relational-database-parameters",
	Short: "Returns all of the runtime parameters offered by the underlying database software, or engine, for a specific database in Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getRelationalDatabaseParametersCmd).Standalone()

	lightsail_getRelationalDatabaseParametersCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	lightsail_getRelationalDatabaseParametersCmd.Flags().String("relational-database-name", "", "The name of your database for which to get parameters.")
	lightsail_getRelationalDatabaseParametersCmd.MarkFlagRequired("relational-database-name")
	lightsailCmd.AddCommand(lightsail_getRelationalDatabaseParametersCmd)
}
