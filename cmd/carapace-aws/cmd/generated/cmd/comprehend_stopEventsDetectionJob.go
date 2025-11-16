package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_stopEventsDetectionJobCmd = &cobra.Command{
	Use:   "stop-events-detection-job",
	Short: "Stops an events detection job in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_stopEventsDetectionJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_stopEventsDetectionJobCmd).Standalone()

		comprehend_stopEventsDetectionJobCmd.Flags().String("job-id", "", "The identifier of the events detection job to stop.")
		comprehend_stopEventsDetectionJobCmd.MarkFlagRequired("job-id")
	})
	comprehendCmd.AddCommand(comprehend_stopEventsDetectionJobCmd)
}
