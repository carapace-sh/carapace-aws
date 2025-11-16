package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_startDominantLanguageDetectionJobCmd = &cobra.Command{
	Use:   "start-dominant-language-detection-job",
	Short: "Starts an asynchronous dominant language detection job for a collection of documents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_startDominantLanguageDetectionJobCmd).Standalone()

	comprehend_startDominantLanguageDetectionJobCmd.Flags().String("client-request-token", "", "A unique identifier for the request.")
	comprehend_startDominantLanguageDetectionJobCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend read access to your input data.")
	comprehend_startDominantLanguageDetectionJobCmd.Flags().String("input-data-config", "", "Specifies the format and location of the input data for the job.")
	comprehend_startDominantLanguageDetectionJobCmd.Flags().String("job-name", "", "An identifier for the job.")
	comprehend_startDominantLanguageDetectionJobCmd.Flags().String("output-data-config", "", "Specifies where to send the output files.")
	comprehend_startDominantLanguageDetectionJobCmd.Flags().String("tags", "", "Tags to associate with the dominant language detection job.")
	comprehend_startDominantLanguageDetectionJobCmd.Flags().String("volume-kms-key-id", "", "ID for the Amazon Web Services Key Management Service (KMS) key that Amazon Comprehend uses to encrypt data on the storage volume attached to the ML compute instance(s) that process the analysis job.")
	comprehend_startDominantLanguageDetectionJobCmd.Flags().String("vpc-config", "", "Configuration parameters for an optional private Virtual Private Cloud (VPC) containing the resources you are using for your dominant language detection job.")
	comprehend_startDominantLanguageDetectionJobCmd.MarkFlagRequired("data-access-role-arn")
	comprehend_startDominantLanguageDetectionJobCmd.MarkFlagRequired("input-data-config")
	comprehend_startDominantLanguageDetectionJobCmd.MarkFlagRequired("output-data-config")
	comprehendCmd.AddCommand(comprehend_startDominantLanguageDetectionJobCmd)
}
