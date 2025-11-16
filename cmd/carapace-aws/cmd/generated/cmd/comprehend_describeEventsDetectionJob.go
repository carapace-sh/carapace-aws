package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_describeEventsDetectionJobCmd = &cobra.Command{
	Use:   "describe-events-detection-job",
	Short: "Gets the status and details of an events detection job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_describeEventsDetectionJobCmd).Standalone()

	comprehend_describeEventsDetectionJobCmd.Flags().String("job-id", "", "The identifier of the events detection job.")
	comprehend_describeEventsDetectionJobCmd.MarkFlagRequired("job-id")
	comprehendCmd.AddCommand(comprehend_describeEventsDetectionJobCmd)
}
