package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_describeApiDestinationCmd = &cobra.Command{
	Use:   "describe-api-destination",
	Short: "Retrieves details about an API destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_describeApiDestinationCmd).Standalone()

	events_describeApiDestinationCmd.Flags().String("name", "", "The name of the API destination to retrieve.")
	events_describeApiDestinationCmd.MarkFlagRequired("name")
	eventsCmd.AddCommand(events_describeApiDestinationCmd)
}
