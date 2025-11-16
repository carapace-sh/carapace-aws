package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_describeSentimentDetectionJobCmd = &cobra.Command{
	Use:   "describe-sentiment-detection-job",
	Short: "Gets the properties associated with a sentiment detection job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_describeSentimentDetectionJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_describeSentimentDetectionJobCmd).Standalone()

		comprehend_describeSentimentDetectionJobCmd.Flags().String("job-id", "", "The identifier that Amazon Comprehend generated for the job.")
		comprehend_describeSentimentDetectionJobCmd.MarkFlagRequired("job-id")
	})
	comprehendCmd.AddCommand(comprehend_describeSentimentDetectionJobCmd)
}
