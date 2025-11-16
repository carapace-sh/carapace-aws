package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_startKeyPhrasesDetectionJobCmd = &cobra.Command{
	Use:   "start-key-phrases-detection-job",
	Short: "Starts an asynchronous key phrase detection job for a collection of documents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_startKeyPhrasesDetectionJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_startKeyPhrasesDetectionJobCmd).Standalone()

		comprehend_startKeyPhrasesDetectionJobCmd.Flags().String("client-request-token", "", "A unique identifier for the request.")
		comprehend_startKeyPhrasesDetectionJobCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend read access to your input data.")
		comprehend_startKeyPhrasesDetectionJobCmd.Flags().String("input-data-config", "", "Specifies the format and location of the input data for the job.")
		comprehend_startKeyPhrasesDetectionJobCmd.Flags().String("job-name", "", "The identifier of the job.")
		comprehend_startKeyPhrasesDetectionJobCmd.Flags().String("language-code", "", "The language of the input documents.")
		comprehend_startKeyPhrasesDetectionJobCmd.Flags().String("output-data-config", "", "Specifies where to send the output files.")
		comprehend_startKeyPhrasesDetectionJobCmd.Flags().String("tags", "", "Tags to associate with the key phrases detection job.")
		comprehend_startKeyPhrasesDetectionJobCmd.Flags().String("volume-kms-key-id", "", "ID for the Amazon Web Services Key Management Service (KMS) key that Amazon Comprehend uses to encrypt data on the storage volume attached to the ML compute instance(s) that process the analysis job.")
		comprehend_startKeyPhrasesDetectionJobCmd.Flags().String("vpc-config", "", "Configuration parameters for an optional private Virtual Private Cloud (VPC) containing the resources you are using for your key phrases detection job.")
		comprehend_startKeyPhrasesDetectionJobCmd.MarkFlagRequired("data-access-role-arn")
		comprehend_startKeyPhrasesDetectionJobCmd.MarkFlagRequired("input-data-config")
		comprehend_startKeyPhrasesDetectionJobCmd.MarkFlagRequired("language-code")
		comprehend_startKeyPhrasesDetectionJobCmd.MarkFlagRequired("output-data-config")
	})
	comprehendCmd.AddCommand(comprehend_startKeyPhrasesDetectionJobCmd)
}
