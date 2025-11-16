package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getRelationalDatabaseBlueprintsCmd = &cobra.Command{
	Use:   "get-relational-database-blueprints",
	Short: "Returns a list of available database blueprints in Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getRelationalDatabaseBlueprintsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getRelationalDatabaseBlueprintsCmd).Standalone()

		lightsail_getRelationalDatabaseBlueprintsCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	})
	lightsailCmd.AddCommand(lightsail_getRelationalDatabaseBlueprintsCmd)
}
