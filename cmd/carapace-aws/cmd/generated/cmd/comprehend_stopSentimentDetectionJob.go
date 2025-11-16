package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_stopSentimentDetectionJobCmd = &cobra.Command{
	Use:   "stop-sentiment-detection-job",
	Short: "Stops a sentiment detection job in progress.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_stopSentimentDetectionJobCmd).Standalone()

	comprehend_stopSentimentDetectionJobCmd.Flags().String("job-id", "", "The identifier of the sentiment detection job to stop.")
	comprehend_stopSentimentDetectionJobCmd.MarkFlagRequired("job-id")
	comprehendCmd.AddCommand(comprehend_stopSentimentDetectionJobCmd)
}
