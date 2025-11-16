package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_deleteApiDestinationCmd = &cobra.Command{
	Use:   "delete-api-destination",
	Short: "Deletes the specified API destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_deleteApiDestinationCmd).Standalone()

	events_deleteApiDestinationCmd.Flags().String("name", "", "The name of the destination to delete.")
	events_deleteApiDestinationCmd.MarkFlagRequired("name")
	eventsCmd.AddCommand(events_deleteApiDestinationCmd)
}
