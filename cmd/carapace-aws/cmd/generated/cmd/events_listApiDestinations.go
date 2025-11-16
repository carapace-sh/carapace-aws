package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_listApiDestinationsCmd = &cobra.Command{
	Use:   "list-api-destinations",
	Short: "Retrieves a list of API destination in the account in the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_listApiDestinationsCmd).Standalone()

	events_listApiDestinationsCmd.Flags().String("connection-arn", "", "The ARN of the connection specified for the API destination.")
	events_listApiDestinationsCmd.Flags().String("limit", "", "The maximum number of API destinations to include in the response.")
	events_listApiDestinationsCmd.Flags().String("name-prefix", "", "A name prefix to filter results returned.")
	events_listApiDestinationsCmd.Flags().String("next-token", "", "The token returned by a previous call, which you can use to retrieve the next set of results.")
	eventsCmd.AddCommand(events_listApiDestinationsCmd)
}
