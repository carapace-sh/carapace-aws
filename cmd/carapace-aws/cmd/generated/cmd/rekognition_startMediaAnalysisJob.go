package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_startMediaAnalysisJobCmd = &cobra.Command{
	Use:   "start-media-analysis-job",
	Short: "Initiates a new media analysis job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_startMediaAnalysisJobCmd).Standalone()

	rekognition_startMediaAnalysisJobCmd.Flags().String("client-request-token", "", "Idempotency token used to prevent the accidental creation of duplicate versions.")
	rekognition_startMediaAnalysisJobCmd.Flags().String("input", "", "Input data to be analyzed by the job.")
	rekognition_startMediaAnalysisJobCmd.Flags().String("job-name", "", "The name of the job.")
	rekognition_startMediaAnalysisJobCmd.Flags().String("kms-key-id", "", "The identifier of customer managed AWS KMS key (name or ARN).")
	rekognition_startMediaAnalysisJobCmd.Flags().String("operations-config", "", "Configuration options for the media analysis job to be created.")
	rekognition_startMediaAnalysisJobCmd.Flags().String("output-config", "", "The Amazon S3 bucket location to store the results.")
	rekognition_startMediaAnalysisJobCmd.MarkFlagRequired("input")
	rekognition_startMediaAnalysisJobCmd.MarkFlagRequired("operations-config")
	rekognition_startMediaAnalysisJobCmd.MarkFlagRequired("output-config")
	rekognitionCmd.AddCommand(rekognition_startMediaAnalysisJobCmd)
}
