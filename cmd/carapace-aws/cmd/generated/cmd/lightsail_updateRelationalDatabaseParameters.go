package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_updateRelationalDatabaseParametersCmd = &cobra.Command{
	Use:   "update-relational-database-parameters",
	Short: "Allows the update of one or more parameters of a database in Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_updateRelationalDatabaseParametersCmd).Standalone()

	lightsail_updateRelationalDatabaseParametersCmd.Flags().String("parameters", "", "The database parameters to update.")
	lightsail_updateRelationalDatabaseParametersCmd.Flags().String("relational-database-name", "", "The name of your database for which to update parameters.")
	lightsail_updateRelationalDatabaseParametersCmd.MarkFlagRequired("parameters")
	lightsail_updateRelationalDatabaseParametersCmd.MarkFlagRequired("relational-database-name")
	lightsailCmd.AddCommand(lightsail_updateRelationalDatabaseParametersCmd)
}
