package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_putProjectEventsCmd = &cobra.Command{
	Use:   "put-project-events",
	Short: "Sends performance events to Evidently.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_putProjectEventsCmd).Standalone()

	evidently_putProjectEventsCmd.Flags().String("events", "", "An array of event structures that contain the performance data that is being sent to Evidently.")
	evidently_putProjectEventsCmd.Flags().String("project", "", "The name or ARN of the project to write the events to.")
	evidently_putProjectEventsCmd.MarkFlagRequired("events")
	evidently_putProjectEventsCmd.MarkFlagRequired("project")
	evidentlyCmd.AddCommand(evidently_putProjectEventsCmd)
}
