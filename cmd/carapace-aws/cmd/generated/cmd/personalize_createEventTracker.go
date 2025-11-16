package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_createEventTrackerCmd = &cobra.Command{
	Use:   "create-event-tracker",
	Short: "Creates an event tracker that you use when adding event data to a specified dataset group using the [PutEvents](https://docs.aws.amazon.com/personalize/latest/dg/API_UBS_PutEvents.html) API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_createEventTrackerCmd).Standalone()

	personalize_createEventTrackerCmd.Flags().String("dataset-group-arn", "", "The Amazon Resource Name (ARN) of the dataset group that receives the event data.")
	personalize_createEventTrackerCmd.Flags().String("name", "", "The name for the event tracker.")
	personalize_createEventTrackerCmd.Flags().String("tags", "", "A list of [tags](https://docs.aws.amazon.com/personalize/latest/dg/tagging-resources.html) to apply to the event tracker.")
	personalize_createEventTrackerCmd.MarkFlagRequired("dataset-group-arn")
	personalize_createEventTrackerCmd.MarkFlagRequired("name")
	personalizeCmd.AddCommand(personalize_createEventTrackerCmd)
}
