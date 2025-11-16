package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_deleteEventTrackerCmd = &cobra.Command{
	Use:   "delete-event-tracker",
	Short: "Deletes the event tracker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_deleteEventTrackerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_deleteEventTrackerCmd).Standalone()

		personalize_deleteEventTrackerCmd.Flags().String("event-tracker-arn", "", "The Amazon Resource Name (ARN) of the event tracker to delete.")
		personalize_deleteEventTrackerCmd.MarkFlagRequired("event-tracker-arn")
	})
	personalizeCmd.AddCommand(personalize_deleteEventTrackerCmd)
}
