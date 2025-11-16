package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_describeEventTrackerCmd = &cobra.Command{
	Use:   "describe-event-tracker",
	Short: "Describes an event tracker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_describeEventTrackerCmd).Standalone()

	personalize_describeEventTrackerCmd.Flags().String("event-tracker-arn", "", "The Amazon Resource Name (ARN) of the event tracker to describe.")
	personalize_describeEventTrackerCmd.MarkFlagRequired("event-tracker-arn")
	personalizeCmd.AddCommand(personalize_describeEventTrackerCmd)
}
