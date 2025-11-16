package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_stopTargetedSentimentDetectionJobCmd = &cobra.Command{
	Use:   "stop-targeted-sentiment-detection-job",
	Short: "Stops a targeted sentiment detection job in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_stopTargetedSentimentDetectionJobCmd).Standalone()

	comprehend_stopTargetedSentimentDetectionJobCmd.Flags().String("job-id", "", "The identifier of the targeted sentiment detection job to stop.")
	comprehend_stopTargetedSentimentDetectionJobCmd.MarkFlagRequired("job-id")
	comprehendCmd.AddCommand(comprehend_stopTargetedSentimentDetectionJobCmd)
}
