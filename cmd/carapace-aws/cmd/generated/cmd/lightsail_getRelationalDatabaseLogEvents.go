package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getRelationalDatabaseLogEventsCmd = &cobra.Command{
	Use:   "get-relational-database-log-events",
	Short: "Returns a list of log events for a database in Amazon Lightsail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getRelationalDatabaseLogEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getRelationalDatabaseLogEventsCmd).Standalone()

		lightsail_getRelationalDatabaseLogEventsCmd.Flags().String("end-time", "", "The end of the time interval from which to get log events.")
		lightsail_getRelationalDatabaseLogEventsCmd.Flags().String("log-stream-name", "", "The name of the log stream.")
		lightsail_getRelationalDatabaseLogEventsCmd.Flags().String("page-token", "", "The token to advance to the next or previous page of results from your request.")
		lightsail_getRelationalDatabaseLogEventsCmd.Flags().String("relational-database-name", "", "The name of your database for which to get log events.")
		lightsail_getRelationalDatabaseLogEventsCmd.Flags().String("start-from-head", "", "Parameter to specify if the log should start from head or tail.")
		lightsail_getRelationalDatabaseLogEventsCmd.Flags().String("start-time", "", "The start of the time interval from which to get log events.")
		lightsail_getRelationalDatabaseLogEventsCmd.MarkFlagRequired("log-stream-name")
		lightsail_getRelationalDatabaseLogEventsCmd.MarkFlagRequired("relational-database-name")
	})
	lightsailCmd.AddCommand(lightsail_getRelationalDatabaseLogEventsCmd)
}
