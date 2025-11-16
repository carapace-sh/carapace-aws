package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getRelationalDatabaseEventsCmd = &cobra.Command{
	Use:   "get-relational-database-events",
	Short: "Returns a list of events for a specific database in Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getRelationalDatabaseEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getRelationalDatabaseEventsCmd).Standalone()

		lightsail_getRelationalDatabaseEventsCmd.Flags().String("duration-in-minutes", "", "The number of minutes in the past from which to retrieve events.")
		lightsail_getRelationalDatabaseEventsCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
		lightsail_getRelationalDatabaseEventsCmd.Flags().String("relational-database-name", "", "The name of the database from which to get events.")
		lightsail_getRelationalDatabaseEventsCmd.MarkFlagRequired("relational-database-name")
	})
	lightsailCmd.AddCommand(lightsail_getRelationalDatabaseEventsCmd)
}
