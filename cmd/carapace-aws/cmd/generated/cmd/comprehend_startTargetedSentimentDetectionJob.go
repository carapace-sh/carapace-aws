package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_startTargetedSentimentDetectionJobCmd = &cobra.Command{
	Use:   "start-targeted-sentiment-detection-job",
	Short: "Starts an asynchronous targeted sentiment detection job for a collection of documents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_startTargetedSentimentDetectionJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_startTargetedSentimentDetectionJobCmd).Standalone()

		comprehend_startTargetedSentimentDetectionJobCmd.Flags().String("client-request-token", "", "A unique identifier for the request.")
		comprehend_startTargetedSentimentDetectionJobCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend read access to your input data.")
		comprehend_startTargetedSentimentDetectionJobCmd.Flags().String("input-data-config", "", "")
		comprehend_startTargetedSentimentDetectionJobCmd.Flags().String("job-name", "", "The identifier of the job.")
		comprehend_startTargetedSentimentDetectionJobCmd.Flags().String("language-code", "", "The language of the input documents.")
		comprehend_startTargetedSentimentDetectionJobCmd.Flags().String("output-data-config", "", "Specifies where to send the output files.")
		comprehend_startTargetedSentimentDetectionJobCmd.Flags().String("tags", "", "Tags to associate with the targeted sentiment detection job.")
		comprehend_startTargetedSentimentDetectionJobCmd.Flags().String("volume-kms-key-id", "", "ID for the KMS key that Amazon Comprehend uses to encrypt data on the storage volume attached to the ML compute instance(s) that process the analysis job.")
		comprehend_startTargetedSentimentDetectionJobCmd.Flags().String("vpc-config", "", "")
		comprehend_startTargetedSentimentDetectionJobCmd.MarkFlagRequired("data-access-role-arn")
		comprehend_startTargetedSentimentDetectionJobCmd.MarkFlagRequired("input-data-config")
		comprehend_startTargetedSentimentDetectionJobCmd.MarkFlagRequired("language-code")
		comprehend_startTargetedSentimentDetectionJobCmd.MarkFlagRequired("output-data-config")
	})
	comprehendCmd.AddCommand(comprehend_startTargetedSentimentDetectionJobCmd)
}
