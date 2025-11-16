package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_startSentimentDetectionJobCmd = &cobra.Command{
	Use:   "start-sentiment-detection-job",
	Short: "Starts an asynchronous sentiment detection job for a collection of documents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_startSentimentDetectionJobCmd).Standalone()

	comprehend_startSentimentDetectionJobCmd.Flags().String("client-request-token", "", "A unique identifier for the request.")
	comprehend_startSentimentDetectionJobCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend read access to your input data.")
	comprehend_startSentimentDetectionJobCmd.Flags().String("input-data-config", "", "Specifies the format and location of the input data for the job.")
	comprehend_startSentimentDetectionJobCmd.Flags().String("job-name", "", "The identifier of the job.")
	comprehend_startSentimentDetectionJobCmd.Flags().String("language-code", "", "The language of the input documents.")
	comprehend_startSentimentDetectionJobCmd.Flags().String("output-data-config", "", "Specifies where to send the output files.")
	comprehend_startSentimentDetectionJobCmd.Flags().String("tags", "", "Tags to associate with the sentiment detection job.")
	comprehend_startSentimentDetectionJobCmd.Flags().String("volume-kms-key-id", "", "ID for the Amazon Web Services Key Management Service (KMS) key that Amazon Comprehend uses to encrypt data on the storage volume attached to the ML compute instance(s) that process the analysis job.")
	comprehend_startSentimentDetectionJobCmd.Flags().String("vpc-config", "", "Configuration parameters for an optional private Virtual Private Cloud (VPC) containing the resources you are using for your sentiment detection job.")
	comprehend_startSentimentDetectionJobCmd.MarkFlagRequired("data-access-role-arn")
	comprehend_startSentimentDetectionJobCmd.MarkFlagRequired("input-data-config")
	comprehend_startSentimentDetectionJobCmd.MarkFlagRequired("language-code")
	comprehend_startSentimentDetectionJobCmd.MarkFlagRequired("output-data-config")
	comprehendCmd.AddCommand(comprehend_startSentimentDetectionJobCmd)
}
