package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_describeTargetedSentimentDetectionJobCmd = &cobra.Command{
	Use:   "describe-targeted-sentiment-detection-job",
	Short: "Gets the properties associated with a targeted sentiment detection job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_describeTargetedSentimentDetectionJobCmd).Standalone()

	comprehend_describeTargetedSentimentDetectionJobCmd.Flags().String("job-id", "", "The identifier that Amazon Comprehend generated for the job.")
	comprehend_describeTargetedSentimentDetectionJobCmd.MarkFlagRequired("job-id")
	comprehendCmd.AddCommand(comprehend_describeTargetedSentimentDetectionJobCmd)
}
